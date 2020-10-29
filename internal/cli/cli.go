package cli

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/mvisonneau/mmds/internal/cmd"
)

// Run handles the instanciation of the CLI application
func Run(version string, args []string) {
	err := NewApp(version, time.Now()).Run(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// NewApp configures the CLI application
func NewApp(version string, start time.Time) (app *cli.App) {
	app = cli.NewApp()
	app.Name = "mmds"
	app.Version = version
	app.Usage = "Missed (AWS) Meta-Data (service)"
	app.EnableBashCompletion = true

	app.Flags = cli.FlagsByName{
		&cli.StringFlag{
			Name:    "log-level",
			EnvVars: []string{"MMDS_LOG_LEVEL"},
			Usage:   "log `level` (debug,info,warn,fatal,panic)",
			Value:   "info",
		},
		&cli.StringFlag{
			Name:    "log-format",
			EnvVars: []string{"MMDS_LOG_FORMAT"},
			Usage:   "log `format` (json,text)",
			Value:   "text",
		},
	}

	app.Commands = cli.CommandsByName{
		{
			Name:   "pricing-model",
			Usage:  "get instance pricing-model",
			Action: cmd.ExecWrapper(cmd.PricingModel),
		},
	}

	app.Metadata = map[string]interface{}{
		"startTime": start,
	}

	return
}
