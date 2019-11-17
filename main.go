package main

import (
	"capybara-heartbeat/cmd"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

func main() {
	// @todo - load the generated config file.
	// @todo - if no config file loaded, then run the auth command first as we need to authenticate before running.
	// @todo - if there's a generated config file, then we can just run with that

	app := &cli.App{
		Name:        "Capybara Heartbeat",
		HelpName:    "",
		Usage:       "https://capybara.dev",
		UsageText:   "",
		ArgsUsage:   "",
		Version:     "0.0.1",
		Description: "",
		Compiled:    time.Time{},
		Commands: []*cli.Command{
			&cli.Command{
				Name:            "run",
				Usage:           "",
				UsageText:       "",
				Description:     "Runs the stats monitoring service.",
				SkipFlagParsing: false,
				HideHelp:        false,
				Hidden:          false,
				HelpName:        "run",
				Action:          cmd.Run,
			},
			&cli.Command{
				Name:        "auth",
				Usage:       `--key "my_key"`,
				UsageText:   "auth - provide your auth key",
				Description: "Authorizes with your Capybara.dev account!",
				ArgsUsage:   "[key]",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "key", Aliases: []string{"k"}},
				},
				SkipFlagParsing: false,
				HideHelp:        false,
				Hidden:          false,
				HelpName:        "auth",
				Action:          cmd.Auth,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
