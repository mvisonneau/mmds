package cmd

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/mvisonneau/go-helpers/logger"
	"github.com/mvisonneau/mmds/pkg/mmds"
)

var start time.Time

func configure(ctx *cli.Context) (i mmds.Instance, err error) {
	start = ctx.App.Metadata["startTime"].(time.Time)

	if err = logger.Configure(logger.Config{
		Level:  ctx.String("log-level"),
		Format: ctx.String("log-format"),
	}); err != nil {
		return
	}

	return mmds.New()
}

func exit(exitCode int, err error) cli.ExitCoder {
	defer log.WithFields(
		log.Fields{
			"execution-time": time.Since(start),
		},
	).Debug("exited..")

	if err != nil {
		log.Error(err.Error())
	}

	return cli.NewExitError("", exitCode)
}

// ExecWrapper ..
func ExecWrapper(f func(ctx *cli.Context) (int, error)) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		return exit(f(ctx))
	}
}
