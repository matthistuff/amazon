package actions

import (
	"github.com/codegangsta/cli"
	"github.com/matthistuff/amazon/config"
	"github.com/matthistuff/amazon/api"
	"github.com/matthistuff/amazon/data"
	"encoding/xml"
	"fmt"
)

func Info(c *cli.Context) {
	api := api.Create(c.GlobalString("locale"))
	index := c.Args().First()
	config := config.GetConfig()
	asin := config.ASINFromCache("Products", index)

	result, err := api.ProductAPI.ItemLookup(asin)
	if err != nil {
		panic(err)
		return
	}

	var lookupResult data.ItemLookupResponse
	if err := xml.Unmarshal([]byte(result), &lookupResult); err != nil {
		panic(err)
		return
	}

	fmt.Println(result)
}
