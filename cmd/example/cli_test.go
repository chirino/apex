package test

import (
	"github.com/urfave/cli/v2"
	"testing"
)

func TestCommonFlags(t *testing.T) {

	commonFlags := []cli.Flag{
		&cli.StringFlag{
			Name: "result",
		},
	}

	result := ""
	app := &cli.App{
		Name:  "root",
		Flags: commonFlags,
		Commands: []*cli.Command{
			{
				Name:  "sub",
				Flags: commonFlags,
				Action: func(ctx *cli.Context) error {
					result = ctx.String("result")
					return nil
				},
			},
		},
	}

	if err := app.Run([]string{"root", "sub", "--result", "after"}); err != nil {
		t.Fatal(err)
	}
	if result != "after" {
		t.Fatalf("result is not after")
	}

	if err := app.Run([]string{"root", "--result", "before", "sub"}); err != nil {
		t.Fatal(err)
	}
	if result != "before" {
		t.Fatalf("result is not before")
	}

}
