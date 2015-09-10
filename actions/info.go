package actions

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/matthistuff/amazon/api"
	"github.com/matthistuff/amazon/color"
	"github.com/matthistuff/amazon/config"
	"github.com/matthistuff/amazon/helper"
	"os"
	"strings"
)

// Info prints information about a product
func Info(c *cli.Context) {
	color.Allow(c)

	api := api.Create(c.GlobalString("locale"))
	config := config.GetConfig()

	asin, exists := config.ASINFromCache("Products", c.Args().First())
	if !exists {
		fmt.Fprintln(os.Stderr, "Cannot look up ASIN")
		os.Exit(1)
	}

	result, err := api.ItemLookup(asin, "ItemAttributes,Small,OfferFull,EditorialReview")
	if err != nil {
		panic(err)
	}

	item := result.Items.ItemList[0]
	rating, err := helper.Rating(item.ASIN, api.Locale)
	if err != nil {
		panic(err)
	}

	year := item.ItemAttributes.PublicationDate
	if year == "" {
		year = item.ItemAttributes.ReleaseDate
	}
	if year != "" {
		year = fmt.Sprintf(" - %4.4s", year)
	}

	fmt.Printf("%s %s\n", color.Header(item.ItemAttributes.Title), helper.FormatRating(rating))

	fmt.Println(color.Faint(fmt.Sprintf("(%s) %s%s",
		strings.Join(item.ItemAttributes.Languages.Languages(), ", "),
		item.ItemAttributes.Binding,
		year)))

	fmt.Println("")
	if item.ItemAttributes.ListPrice.FormattedPrice != "" {
		fmt.Printf("%s\n", color.Bold(item.ItemAttributes.ListPrice.FormattedPrice))
	}
	if item.OfferSummary.TotalNew > 0 {
		fmt.Printf("%d new from %s\t", item.OfferSummary.TotalNew, item.OfferSummary.LowestNewPrice.FormattedPrice)
	}
	if item.OfferSummary.TotalUsed > 0 {
		fmt.Printf("%d used from %s\t", item.OfferSummary.TotalUsed, item.OfferSummary.LowestUsedPrice.FormattedPrice)
	}
	fmt.Printf("\n\n")

	if len(item.ItemAttributes.Authors) > 0 {
		fmt.Println(color.Header("AUTHORS"))
		for _, author := range item.ItemAttributes.Authors {
			fmt.Printf("\t%s\n", author)
		}
	}

	if item.ItemAttributes.Publisher != "" {
		fmt.Printf("%s\n\t%s\n", color.Header("PUBLISHER"), item.ItemAttributes.Publisher)
	}

	if item.ItemAttributes.Edition != "" {
		fmt.Printf("%s\n\t%s\n", color.Header("EDITION"), item.ItemAttributes.Edition)
	}

	if item.ItemAttributes.ISBN != "" {
		fmt.Printf("%s\n\t%s\n", color.Header("ISBN"), item.ItemAttributes.ISBN)
	}

	if len(item.EditorialReviews.EditorialReviewList) > 0 {
		fmt.Printf("%s\n\t%s\n",
			color.Header(strings.ToUpper(item.EditorialReviews.EditorialReviewList[0].Source)),
			item.EditorialReviews.EditorialReviewList[0].Content)
	}
}
