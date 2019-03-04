package main

import (
	"github.com/urfave/cli"
)

var version = "<devel>"

// runCli : Generates cli configuration for the application
func runCli() (c *cli.App) {
	c = cli.NewApp()
	c.Name = "mmds"
	c.Version = version
	c.Usage = "Missed (AWS) Meta-Data (service)"
	c.EnableBashCompletion = true

	c.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "log-level",
			EnvVar:      "MMDS_LOG_LEVEL",
			Usage:       "log `level` (debug,info,warn,fatal,panic)",
			Value:       "info",
			Destination: &cfg.Log.Level,
		},
		cli.StringFlag{
			Name:        "log-format",
			EnvVar:      "MMDS_LOG_FORMAT",
			Usage:       "log `format` (json,text)",
			Value:       "text",
			Destination: &cfg.Log.Format,
		},
	}

	c.Commands = []cli.Command{
		{
			Name:   "pricing-model",
			Usage:  "get instance pricing-model",
			Action: run,
		},
	}

	return
}
