package actions

import (
	"github.com/codegangsta/cli"
	"github.com/matthistuff/amazon/config"
	"github.com/matthistuff/amazon/api"
	"github.com/pkg/browser"
	"fmt"
	"os"
)

func Open(c *cli.Context) {
	api := api.Create(c.GlobalString("locale"))

	conf := config.GetConfig()
	asin, exists := conf.ASINFromCache("Products", c.Args().First())
	if !exists {
		fmt.Errorf("Cannot look up ASIN")
		os.Exit(1)
	}

	result, err := api.ItemLookup(asin)

	if err != nil {
		panic(err)
		return
	}

	browser.OpenURL(result.Items.ItemList[0].DetailPageURL)
}
