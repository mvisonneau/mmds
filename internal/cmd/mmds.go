package cmd

import (
	"fmt"

	"github.com/mvisonneau/mmds/pkg/mmds"
	"github.com/urfave/cli/v2"
)

// PricingModel ..
func PricingModel(ctx *cli.Context) (exit int, err error) {
	var i mmds.Instance
	if i, err = configure(ctx); err != nil {
		exit = 1
		return
	}

	fmt.Println(i.GetPricingModel())
	return
}
