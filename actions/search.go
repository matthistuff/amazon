package actions

import (
	"fmt"
	"github.com/codegangsta/cli"
	"strings"
	"github.com/matthistuff/amazon/api"
	"strconv"
	"github.com/matthistuff/amazon/config"
	"github.com/matthistuff/amazon/color"
)

func Search(c *cli.Context) {
	color.Allow(c)

	search := strings.Replace(strings.Join(c.Args(), "+"), " ", "+", -1)
	page := c.Int("page")
	api := api.Create(c.GlobalString("locale"))

	conf := config.GetConfig()
	defer conf.Flush()

	params := map[string]string{
		"Keywords":      search,
		"ResponseGroup": "ItemAttributes,Small,EditorialReview,OfferSummary,BrowseNodes",
		"ItemPage":      strconv.FormatInt(int64(page), 10),
	}
	result, err := api.ItemSearch("All", params)

	if err != nil {
		panic(err)
		return
	}

	fmt.Printf("---\nFound %d results matching query %s\n---\n", result.Items.TotalResults, color.Header("'%s'", search))

	cache := make(map[string]string)
	for index, item := range result.Items.ItemList {
		price := item.ItemAttributes.ListPrice.FormattedPrice

		if price == "" {
			if lowestNew := item.OfferSummary.LowestNewPrice; lowestNew.Amount != 0 {
				price = fmt.Sprintf("%s (new)", lowestNew.FormattedPrice)
			}
		}

		normalizedIndex := index + 1
		cache[strconv.Itoa(normalizedIndex)] = item.ASIN
		fmt.Printf("(%s) %-45.45s %18s [%s]\n", color.ShortId("%2d", normalizedIndex), item.ItemAttributes.Title, color.Bold(price), item.ItemAttributes.Binding)
	}
	conf.ResultCache["Search"] = cache

	fmt.Printf("---\nPage %d of %d\n---\n", page, result.Items.TotalPages)
}