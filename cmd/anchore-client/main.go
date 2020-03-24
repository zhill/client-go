package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "anchore-ctl",
		Version: "0.1.0",
		Description: "Simple CLI for Anchore Engine API",
		Commands: []*cli.Command{
			ImageCommand,
		},
		Flags:                  []cli.Flag{
			&cli.StringFlag{
				Name:        "username",
				Aliases:     []string{"u"},
				EnvVars:     []string{"ANCHORE_CLI_USER"},
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "password",
				Aliases:     []string{"p"},
				EnvVars:     []string{"ANCHORE_CLI_PASS"},
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "url",
				Aliases:     []string{"l"},
				EnvVars:     []string{"ANCHORE_CLI_URL"},
				Required:    true,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("Error. Exiting %v", err.Error())
		os.Exit(1)
	}
}
