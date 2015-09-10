package actions

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/matthistuff/amazon/api"
	"github.com/matthistuff/amazon/config"
	"github.com/pkg/browser"
	"os"
)

// Open opens a product on the amazon website
func Open(c *cli.Context) {
	api := api.Create(c.GlobalString("locale"))

	conf := config.GetConfig()
	asin, exists := conf.ASINFromCache("Products", c.Args().First())
	if !exists {
		fmt.Fprintln(os.Stderr, "Cannot look up ASIN")
		os.Exit(1)
	}

	result, err := api.ItemLookup(asin, "Small")

	if err != nil {
		panic(err)
	}

	browser.OpenURL(result.Items.ItemList[0].DetailPageURL)
}
