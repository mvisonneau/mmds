package main

import (
	"os"

	"github.com/mvisonneau/mmds/internal/cli"
)

var version = ""

func main() {
	cli.Run(version, os.Args)
}
