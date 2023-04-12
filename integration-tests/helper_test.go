package integration_tests

import (
	"context"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/nexodus-io/nexodus/internal/nexodus"
	"github.com/nexodus-io/nexodus/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
	"io"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"
	"unicode"
)

type Helper struct {
	*testing.T
	require *require.Assertions
	assert  *assert.Assertions
	logger  *zap.SugaredLogger
}

func NewHelper(t *testing.T) *Helper {
	return &Helper{
		T:       t,
		require: require.New(t),
		assert:  assert.New(t),
	}
}

func callerSourceLine() string {
	if _, helperFile, _, ok := runtime.Caller(0); ok {
		for skip := 1; skip < 20; skip++ {
			if _, file, line, ok := runtime.Caller(skip); ok {
				if file != helperFile {
					return fmt.Sprintf("%s:%d:", filepath.Base(file), line)
				}
			} else {
				break
			}
		}
	}
	return ""
}

func (helper *Helper) log(args ...any) {
	caller := callerSourceLine()
	if caller != "" {
		args = append([]any{caller}, args...)
	}
	helper.T.Log(args...)
}
func (helper *Helper) logf(fmt string, args ...any) {
	caller := callerSourceLine()
	if caller != "" {
		args = append([]any{caller}, args...)
		helper.T.Logf("%s "+fmt, args...)
	} else {
		helper.T.Logf(fmt, args...)
	}
}

// CreateNode creates a container
func (helper *Helper) CreateNode(ctx context.Context, nameSuffix string, networks []string, v6 v6Enable) (testcontainers.Container, func()) {

	// Host modifiers differ for a container for a container with and without v6 enabled
	var hostConfSysctl map[string]string
	if v6 == enableV6 {
		hostConfSysctl = map[string]string{
			"net.ipv6.conf.all.disable_ipv6": "0",
			"net.ipv4.ip_forward":            "1",
			"net.ipv6.conf.all.forwarding":   "1",
		}
	} else {
		hostConfSysctl = map[string]string{
			"net.ipv4.ip_forward":          "1",
			"net.ipv6.conf.all.forwarding": "1",
		}
	}

	// Name containers <test>-<nameSuffix>, where <test> is the name of the calling function
	name := nameSuffix
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		helper.log("runtime.Caller: failed")
	} else {
		callerParts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
		name = fmt.Sprintf("%s-%s", callerParts[len(callerParts)-1], nameSuffix)
	}

	certsDir, err := findCertsDir()
	require.NoError(helper.T, err)

	req := testcontainers.ContainerRequest{
		Image:    "quay.io/nexodus/nexd:latest",
		Name:     name,
		Networks: networks,
		HostConfigModifier: func(hostConfig *container.HostConfig) {
			hostConfig.Sysctls = hostConfSysctl
			hostConfig.CapAdd = []string{
				"SYS_MODULE",
				"NET_ADMIN",
				"NET_RAW",
			}
			hostConfig.ExtraHosts = []string{
				fmt.Sprintf("try.nexodus.127.0.0.1.nip.io:%s", hostDNSName),
				fmt.Sprintf("api.try.nexodus.127.0.0.1.nip.io:%s", hostDNSName),
				fmt.Sprintf("auth.try.nexodus.127.0.0.1.nip.io:%s", hostDNSName),
			}
			hostConfig.AutoRemove = true
		},
		Mounts: []testcontainers.ContainerMount{
			{
				Source: testcontainers.GenericBindMountSource{
					HostPath: certsDir,
				},
				Target:   "/.certs",
				ReadOnly: true,
			},
		},
		Cmd: []string{
			"/update-ca.sh",
			"/bin/bash",
			"-c",
			"echo ready && sleep infinity",
		},
	}
	ctr, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ProviderType:     providerType,
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(helper.T, err)
	stop := func() {
		if helper.T.Failed() {
			helper.Log(helper.gatherFail(ctr))
		}
		err = ctr.Terminate(context.Background())
		if err != nil {
			helper.logf("%s: container terminate failed: %v", name, err)
		}
	}

	// wait for the CA cert to get imported.
	wg := sync.WaitGroup{}
	wg.Add(1)
	ctr.FollowOutput(FnConsumer{
		Apply: func(l testcontainers.Log) {
			text := strings.TrimRightFunc(string(l.Content), unicode.IsSpace)
			if text == "ready" {
				wg.Done()
				ctr.StopLogProducer()
			}
		},
	})
	ctr.StartLogProducer(ctx)
	wg.Wait()
	return ctr, stop
}

// containerExec exec container commands
func (helper *Helper) containerExec(ctx context.Context, container testcontainers.Container, cmd []string) (string, error) {
	nodeName, _ := container.Name(ctx)
	if cmd[0] != "wg" && cmd[0] != "cat" {
		helper.logf("Running command on %s: %s", nodeName, strings.Join(cmd, " "))
	}
	code, outputRaw, err := container.Exec(
		ctx,
		cmd,
	)
	if err != nil {
		return "", err
	}
	output, err := io.ReadAll(outputRaw)
	if err != nil {
		return "", err
	}
	if code != 0 {
		return "", fmt.Errorf("exit code %d. stderr: %s", code, string(output))
	}

	return string(output), err
}

// CreateNetwork creates a docker network
func (helper *Helper) CreateNetwork(ctx context.Context, name, cidr string) testcontainers.Network {
	req := testcontainers.GenericNetworkRequest{
		ProviderType: providerType,
		NetworkRequest: testcontainers.NetworkRequest{
			Name:   name,
			Driver: "bridge",
			IPAM: &network.IPAM{
				Driver: ipamDriver,
				Config: []network.IPAMConfig{
					{
						Subnet: cidr,
					},
				},
			},
		},
	}
	net, err := testcontainers.GenericNetwork(
		ctx,
		req,
	)
	require.NoError(helper.T, err)
	return net
}

// runNexd copies the nexd command to a file on the container and runs it piping the logs to a file
func (helper *Helper) runNexd(ctx context.Context, node testcontainers.Container, args ...string) {
	nodeName, _ := node.Name(ctx)
	runScript := fmt.Sprintf("%s-nexd-run.sh", strings.TrimPrefix(nodeName, "/"))
	runScriptLocal := fmt.Sprintf("tmp/%s", runScript)
	cmd := []string{"NEXD_LOGLEVEL=debug", "/bin/nexd"}
	cmd = append(cmd, args...)
	cmd = append(cmd, "https://try.nexodus.127.0.0.1.nip.io")
	cmd = append(cmd, ">> /nexd.logs 2>&1 &")

	// write the nexd run command to a local file
	logger := zaptest.NewLogger(helper.T).Sugar()
	nexodus.WriteToFile(logger, strings.Join(cmd, " "), runScriptLocal, 0755)
	// copy the nexd run script to the test container
	err := node.CopyFileToContainer(ctx, runScriptLocal, fmt.Sprintf("/bin/%s", runScript), 0755)
	helper.require.NoError(err, fmt.Errorf("execution of copy command on %s failed: %v", nodeName, err))

	// execute the nexd run script on the test container
	_, err = helper.containerExec(ctx, node, []string{"/bin/bash", "-c", runScript})
	helper.require.NoError(err)
}

// wgDump dump wg sessions for failed test debugging
func (helper *Helper) wgDump(ctx context.Context, container testcontainers.Container) string {
	wgDump, err := helper.containerExec(ctx, container, []string{"wg", "show", "wg0", "dump"})
	if err != nil {
		return ""
	}

	return wgDump
}

// routesDumpV4 dump v4 routes for failed test debugging
func (helper *Helper) routesDumpV4(ctx context.Context, container testcontainers.Container) string {
	routesDump, err := helper.containerExec(ctx, container, []string{"ip", "route"})
	if err != nil {
		return ""
	}

	return routesDump
}

// routesDumpV6 dump v6 routes for failed test debugging
func (helper *Helper) routesDumpV6(ctx context.Context, container testcontainers.Container) string {
	routesDump, err := helper.containerExec(ctx, container, []string{"ip", "-6", "route"})
	if err != nil {
		return ""
	}

	return routesDump
}

// logsDump dump routes for failed test debugging
func (helper *Helper) logsDump(ctx context.Context, container testcontainers.Container) string {
	logsDump, err := helper.containerExec(ctx, container, []string{"cat", "/nexd.logs"})
	if err != nil {
		return "no logs found"
	}

	return logsDump
}

// gatherFail gather details on a failed test for debugging
func (helper *Helper) gatherFail(c testcontainers.Container) string {
	var res []string

	ctx := context.Background()
	nodeName, _ := c.Name(ctx)

	ip, err := getContainerIfaceIPNoRetry(ctx, inetV4, "wg0", c)
	if err != nil {
		helper.log(err)
	}
	res = append(res, fmt.Sprintf("%s wg0 IPv4: %s", nodeName, ip))
	ip, err = getContainerIfaceIPNoRetry(ctx, inetV6, "wg0", c)
	if err != nil {
		helper.log(err)
	}
	res = append(res, fmt.Sprintf("%s wg0 IPv6: %s", nodeName, ip))
	ip, err = getContainerIfaceIPNoRetry(ctx, inetV4, "eth0", c)
	if err != nil {
		helper.log(err)
	}
	res = append(res, fmt.Sprintf("%s eth0 IPv4: %s", nodeName, ip))
	ip, err = getContainerIfaceIPNoRetry(ctx, inetV6, "eth0", c)
	if err != nil {
		helper.log(err)
	}
	res = append(res, fmt.Sprintf("%s eth0 IPv6: %s", nodeName, ip))
	res = append(res, fmt.Sprintf("%s wg-dump:\n %s, ", nodeName, helper.wgDump(ctx, c)))
	res = append(res, fmt.Sprintf("%s routes IPv4:\n%s, ", nodeName, helper.routesDumpV4(ctx, c)))
	res = append(res, fmt.Sprintf("%s routes IPv6:\n%s, ", nodeName, helper.routesDumpV6(ctx, c)))
	res = append(res, fmt.Sprintf("%s nexd logs:\n%s\n, ", nodeName, helper.logsDump(ctx, c)))

	return strings.Join(res, "\n")
}

// runCommand runs the cmd and returns the combined stdout and stderr
func (helper *Helper) runCommand(cmd ...string) (string, error) {
	helper.logf("Running command: %s", strings.Join(cmd, " "))
	output, err := exec.Command(cmd[0], cmd[1:]...).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to run %q: %s (%s)", strings.Join(cmd, " "), err, output)
	}

	return string(output), nil
}

// nexdStatus checks for a Running status of the nexd process via nexctl
func (helper *Helper) nexdStatus(ctx context.Context, ctr testcontainers.Container) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second*1000)
	defer cancel()
	running, _ := util.CheckPeriodically(timeoutCtx, time.Second, func() (bool, error) {
		statOut, _ := helper.containerExec(ctx, ctr, []string{"/bin/nexctl", "nexd", "status"})
		return strings.Contains(statOut, "Running"), nil
	})
	if running {
		return nil
	}
	nodeName, _ := ctr.Name(ctx)
	return fmt.Errorf("failed to get a 'Running' status from the nexd process in node: %s", nodeName)
}

func (helper *Helper) createNewUser(ctx context.Context, password string) (string, func()) {
	id, cleanup, err := createNewUserWithName(ctx, "kitteh", password)
	helper.require.NoError(err)
	return id, cleanup
}

// getNodeHostname trims the container ID down to the node hostname
func (helper *Helper) getNodeHostname(ctx context.Context, ctr testcontainers.Container) (string, error) {
	var hostname string
	err := backoff.Retry(func() error {
		cid := ctr.GetContainerID()
		if len(cid) == 12 {
			hostname = ctr.GetContainerID()
		}
		if len(cid) < 12 {
			return fmt.Errorf("invalid container ID: %s", ctr.GetContainerID())
		} else {
			hostname = strings.TrimSpace(cid[:12])
		}
		return nil
	}, backoff.WithContext(backoff.NewConstantBackOff(1*time.Second), ctx))

	return hostname, err
}
