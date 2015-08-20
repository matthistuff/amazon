package actions

import (
	"fmt"
	"github.com/codegangsta/cli"
	"strings"
	"github.com/matthistuff/amazon/api"
	"strconv"
	"github.com/matthistuff/amazon/config"
	"github.com/matthistuff/amazon/color"
	"math"
	"github.com/matthistuff/amazon/helper"
)

func Search(c *cli.Context) {
	color.Allow(c)

	search := strings.Replace(strings.Join(c.Args(), "+"), " ", "+", -1)
	page := c.Int("page")
	sort := c.String("sort")
	searchIndex := c.String("index")
	api := api.Create(c.GlobalString("locale"))

	conf := config.GetConfig()
	defer conf.Flush()

	params := map[string]string{
		"Keywords":      search,
		"ResponseGroup": "ItemAttributes,Small,EditorialReview,OfferSummary,BrowseNodes",
		"ItemPage":      strconv.FormatInt(int64(page), 10),
	}
	if sort != "" {
		params["Sort"] = sort
	}
	result, err := api.ItemSearch(searchIndex, params)

	if err != nil {
		panic(err)
		return
	}

	fmt.Printf("\nFound %d results matching query %s\n\n", result.Items.TotalResults, color.Header("'%s'", search))

	cache := make(map[string]string)
	for index, item := range result.Items.ItemList {
		price := item.ItemAttributes.ListPrice.FormattedPrice

		if price == "" {
			if lowestNew := item.OfferSummary.LowestNewPrice; lowestNew.Amount != 0 {
				price = fmt.Sprintf("%s (new)", lowestNew.FormattedPrice)
			} else if lowestUsed := item.OfferSummary.LowestUsedPrice; lowestUsed.Amount != 0 {
				price = fmt.Sprintf("%s (used)", lowestUsed.FormattedPrice)
			} else {
				price = "n/a"
			}
		}

		year := item.ItemAttributes.PublicationDate
		if year == "" {
			year = item.ItemAttributes.ReleaseDate
		}
		if year != "" {
			year = fmt.Sprintf(" (%4.4s)", year)
		}

		ratingFormatted := ""
		if !c.Bool("no-rating") {
			rating, err := helper.Rating(item.ASIN, api.Locale)
			if err != nil {
				panic(err)
				return
			}

			ratingFormatted = fmt.Sprintf("%-5s ", helper.FormatRating(rating))
		}

		normalizedIndex := index + 1
		cache[strconv.Itoa(normalizedIndex)] = item.ASIN

		maxLen := math.Min(float64(52 - len(year)), float64(len(item.ItemAttributes.Title)))
		fmt.Printf("(%s) %-52s %s%s [%s]\n",
			color.ShortId("%2d", normalizedIndex),
			fmt.Sprintf("%s%s", item.ItemAttributes.Title[:int(maxLen)], year),
			color.Faint(ratingFormatted),
			color.Bold("%9s", price),
			item.ItemAttributes.Binding)
	}
	conf.ResultCache["Products"] = cache

	fmt.Printf("\nPage %d of %d\n\n", page, result.Items.TotalPages)
}