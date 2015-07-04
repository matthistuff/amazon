package actions

import (
	"github.com/codegangsta/cli"
	"github.com/matthistuff/amazon/config"
	"github.com/matthistuff/amazon/api"
	"github.com/pkg/browser"
)

func Open(c *cli.Context) {
	api := api.Create(c.GlobalString("locale"))

	conf := config.GetConfig()
	asin := conf.ASINFromCache("Search", c.Args().First())

	result, err := api.ItemLookup(asin)

	if err != nil {
		panic(err)
		return
	}

	browser.OpenURL(result.Items.ItemList[0].DetailPageURL)
}
