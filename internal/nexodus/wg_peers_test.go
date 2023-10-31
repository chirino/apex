package nexodus

import (
	"net/netip"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/nexodus-io/nexodus/internal/api/public"
)

func TestRebuildPeerConfig(t *testing.T) {
	zLogger, _ := zap.NewDevelopment()
	testLogger := zLogger.Sugar()
	nxBase := &Nexodus{
		org: &public.ModelsOrganization{
			Cidr:   "100.64.0.0/10",
			CidrV6: "200::/64",
		},
		nodeReflexiveAddressIPv4: netip.MustParseAddrPort("1.1.1.1:1234"),
		logger:                   testLogger,
	}
	nxRelay := &Nexodus{
		org:                      nxBase.org,
		relay:                    true,
		nodeReflexiveAddressIPv4: netip.MustParseAddrPort("1.1.1.1:1234"),
		logger:                   testLogger,
	}
	nxSymmetricNAT := &Nexodus{
		org:                      nxBase.org,
		symmetricNat:             true,
		nodeReflexiveAddressIPv4: netip.MustParseAddrPort("1.1.1.1:1234"),
		logger:                   testLogger,
	}

	testCases := []struct {
		// descriptive name of the test case
		name string
		// the parameters set in the Nexodus object reflect the parameters for the local node
		nx               *Nexodus
		peerLocalIP      string
		peerStunIP       string
		peerIsRelay      bool
		peerSymmetricNAT bool
		// we have a healthy relay available
		healthyRelay bool
		// the peering method expected to be chosen based on the local and remote peer parameters
		expectedMethod string
		// the second choice peering method
		secondMethod string
		// the third choice peering method
		thirdMethod string
	}{
		{
			// Ensure we choose direct peering when the reflexive IPs are the same
			name:           "direct peering",
			nx:             nxBase,
			peerLocalIP:    "192.168.10.50:5678",
			peerStunIP:     "1.1.1.1:4321",
			expectedMethod: peeringMethodDirectLocal,
			secondMethod:   peeringMethodReflexive,
			thirdMethod:    peeringMethodDirectLocal,
		},
		{
			// Ensure we choose direct peering when the reflexive IPs are the same and fall back to a relay
			name:           "direct peering",
			nx:             nxBase,
			peerLocalIP:    "192.168.10.50:5678",
			peerStunIP:     "1.1.1.1:4321",
			healthyRelay:   true,
			expectedMethod: peeringMethodDirectLocal,
			secondMethod:   peeringMethodReflexive,
			thirdMethod:    peeringMethodViaRelay,
		},
		{
			// Ensure we choose reflexive peering when the reflexive IPs are different
			name:           "reflexive peering",
			nx:             nxBase,
			peerLocalIP:    "192.168.10.50:5678",
			peerStunIP:     "2.2.2.2:4321",
			expectedMethod: peeringMethodReflexive,
			secondMethod:   peeringMethodReflexive, // our only choice
			thirdMethod:    peeringMethodReflexive, // our only choice
		},
		{
			// Ensure we choose reflexive peering when the reflexive IPs are different
			name:           "reflexive peering",
			nx:             nxBase,
			peerLocalIP:    "192.168.10.50:5678",
			peerStunIP:     "2.2.2.2:4321",
			healthyRelay:   true,
			expectedMethod: peeringMethodReflexive,
			secondMethod:   peeringMethodViaRelay,
			thirdMethod:    peeringMethodViaRelay, // stay with a healthy relay
		},
		{
			// Peer directly with a relay that is behind the same reflexive IP
			name:           "direct peering to relay",
			nx:             nxBase,
			peerLocalIP:    "192.168.10.50:5678",
			peerStunIP:     "1.1.1.1:4321",
			peerIsRelay:    true,
			expectedMethod: peeringMethodRelayPeerDirectLocal,
			secondMethod:   peeringMethodRelayPeer,
			thirdMethod:    peeringMethodRelayPeerDirectLocal,
		},
		{
			// Peer via the reflexive IP of a relay when not behind the same reflexive IP
			name:           "reflexive peering to relay",
			nx:             nxBase,
			peerLocalIP:    "192.168.10.50:5678",
			peerStunIP:     "2.2.2.2:4321",
			peerIsRelay:    true,
			expectedMethod: peeringMethodRelayPeer,
			secondMethod:   peeringMethodRelayPeer, // our only choice
			thirdMethod:    peeringMethodRelayPeer, // our only choice
		},
		{
			// We are the relay on the same network as a peer
			name:           "direct peering from relay",
			nx:             nxRelay,
			peerLocalIP:    "192.168.10.50:5678",
			peerStunIP:     "1.1.1.1:4321",
			expectedMethod: peeringMethodRelaySelfDirectLocal,
			secondMethod:   peeringMethodRelaySelf,
			thirdMethod:    peeringMethodRelaySelfDirectLocal,
		},
		{
			// Ensure we choose reflexive peering when the reflexive IPs are different
			name:           "reflexive peering",
			nx:             nxRelay,
			peerLocalIP:    "192.168.10.50:5678",
			peerStunIP:     "2.2.2.2:4321",
			expectedMethod: peeringMethodRelaySelf,
			secondMethod:   peeringMethodRelaySelf, // our only choice
			thirdMethod:    peeringMethodRelaySelf, // our only choice
		},
		{
			// Use direct peering when behind the same reflexive IP, even if we are also
			// behind symmetric NAT.
			name:           "direct peering behind symmetric NAT",
			nx:             nxSymmetricNAT,
			peerLocalIP:    "192.168.10.50:5678",
			peerStunIP:     "1.1.1.1:4321",
			expectedMethod: peeringMethodDirectLocal,
			secondMethod:   peeringMethodDirectLocal, // our only choice without a relay
			thirdMethod:    peeringMethodDirectLocal, // our only choice without a relay
		},
		{
			// No peering method available when we are behind symmetric NAT and we
			// have no relay available.
			name:           "no peering method when behind symmetric NAT without a relay",
			nx:             nxSymmetricNAT,
			peerLocalIP:    "192.168.10.50:5678",
			peerStunIP:     "2.2.2.2:4321",
			expectedMethod: peeringMethodNone,
			secondMethod:   peeringMethodNone,
			thirdMethod:    peeringMethodNone,
		},
		{
			// Use the relay when we are behind symmetric NAT and we have a relay available
			name:           "use relay when we are behind symmetric NAT",
			nx:             nxSymmetricNAT,
			peerLocalIP:    "192.168.10.50:5678",
			peerStunIP:     "2.2.2.2:4321",
			healthyRelay:   true,
			expectedMethod: peeringMethodViaRelay,
			secondMethod:   peeringMethodViaRelay, // our only choice
			thirdMethod:    peeringMethodViaRelay, // our only choice
		},
		{
			// Use the relay when the peer is behind symmetric NAT, even if we are not
			name:             "use relay when the peer is behind symmetric NAT",
			nx:               nxBase,
			peerLocalIP:      "192.168.10.50:5678",
			peerStunIP:       "2.2.2.2:4321",
			peerSymmetricNAT: true,
			healthyRelay:     true,
			expectedMethod:   peeringMethodViaRelay,
			secondMethod:     peeringMethodViaRelay, // our only choice
			thirdMethod:      peeringMethodViaRelay, // our only choice
		},
	}

	require := require.New(t)

	for _, tcIter := range testCases {
		tc := tcIter
		t.Run(tc.name, func(t *testing.T) {
			d := deviceCacheEntry{
				device: public.ModelsDevice{
					Endpoints: []public.ModelsEndpoint{
						{
							Address: tc.peerLocalIP,
							Source:  "local",
						},
						{
							Address: tc.peerStunIP,
							Source:  "stun",
						},
					},
					PublicKey:    "bacon",
					Relay:        tc.peerIsRelay,
					SymmetricNat: tc.peerSymmetricNAT,
				},
			}
			tc.nx.peeringReset(&d)

			_, chosenMethod, chosenIndex := tc.nx.rebuildPeerConfig(&d, tc.healthyRelay)
			require.Equal(tc.expectedMethod, chosenMethod)

			now := time.Now()

			// Ensure we stick with the chosen method while the peer is healthy.
			// Set that we peered 15 minutes ago and it was healthy 5 seconds later.
			d.peerHealthy = true
			d.peeringMethod = chosenMethod
			d.peeringMethodIndex = chosenIndex
			d.peeringTime = now.Add(-15 * time.Minute)
			d.peerHealthyTime = now.Add(-15*time.Minute + 5*time.Second)
			_, chosenMethod, _ = tc.nx.rebuildPeerConfig(&d, tc.healthyRelay)
			require.Equal(tc.expectedMethod, chosenMethod)

			// Switch to unhealthy, but since this was previously a successful
			// peer configuration, it should keep trying for 3 minutes. We
			// have 1 minute until the deadline to work again.
			d.peerHealthy = false
			d.peerHealthyTime = now.Add(-1*peeringRestoreTimeout + time.Minute)
			_, chosenMethod, _ = tc.nx.rebuildPeerConfig(&d, tc.healthyRelay)
			require.Equal(tc.expectedMethod, chosenMethod)

			// After 3 minutes, we should switch to the next best method.
			// We were last healthy 3 minutes and 5 seconds ago.
			d.peerHealthyTime = now.Add(-1*peeringRestoreTimeout - 5*time.Second)
			_, chosenMethod, chosenIndex = tc.nx.rebuildPeerConfig(&d, tc.healthyRelay)
			require.Equal(tc.secondMethod, chosenMethod)

			// Another recalculation should switch to the third best method.
			d.peeringMethod = chosenMethod
			d.peeringMethodIndex = chosenIndex
			_, chosenMethod, _ = tc.nx.rebuildPeerConfig(&d, tc.healthyRelay)
			require.Equal(tc.thirdMethod, chosenMethod)
		})
	}
}