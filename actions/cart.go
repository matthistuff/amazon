package actions

import (
	"github.com/codegangsta/cli"
	"github.com/matthistuff/amazon/config"
	"time"
	"github.com/matthistuff/amazon/api"
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/matthistuff/amazon/color"
)

func CartAdd(c *cli.Context) {
	cartName := c.GlobalString("name")
	api := api.Create(c.GlobalString("locale"))

	conf := config.GetConfig()
	defer conf.Flush()

	asin := conf.ASINFromCache("Search", c.Args().First())


	if cart, exists := conf.Carts[cartName]; !exists {
		createResponse, createErr := api.CartCreate(map[string]int{
			asin: 1,
		})

		if createErr != nil {
			panic(createErr)
			return
		}

		conf.Carts[cartName] = &config.Cart{
			Name: cartName,
			Created: time.Now(),
			CartId: createResponse.Cart.CartId,
			HMAC: createResponse.Cart.HMAC,
		}
	} else {
		addResponse, addErr := api.CartAdd(cart.CartId, cart.HMAC, map[string]int{
			asin: 1,
		})

		if addErr != nil {
			panic(addErr)
			return
		}

		conf.Carts[cartName].HMAC = addResponse.Cart.HMAC
	}

	fmt.Printf("Added item to cart %s\n", cartName)
}

func CartInfo(c *cli.Context) {
	color.Allow(c)

	cartName := c.GlobalString("name")
	api := api.Create(c.GlobalString("locale"))

	conf := config.GetConfig()
	defer conf.Flush()

	cartName = conf.NumericFromCache("CartsList", cartName)

	if cart, exists := conf.Carts[cartName]; !exists {
		fmt.Fprintf(os.Stderr, "Cart %s is unknown\n", cartName)
		os.Exit(1)
	} else {
		fmt.Printf("---\nCart %s\n---\n", color.Header(cart.Name))

		getResponse, getErr := api.CartGet(cart.CartId, cart.HMAC)

		if getErr != nil {
			panic(getErr)
			return
		}

		index := 1
		cache := make(map[string]string)
		for _, item := range getResponse.Cart.CartItems.CartItemList {
			fmt.Printf("(%s) %-45.45s %18s × %d\n",
				color.ShortId(strconv.Itoa(index)),
				item.Title,
				item.ItemTotal.FormattedPrice,
				item.Quantity)
			cache[strconv.Itoa(index)] = item.CartItemId
			index += 1
		}
		conf.ResultCache["Cart" + strings.Title(cartName) + "Info"] = cache

		if len(getResponse.Cart.CartItems.CartItemList) == 0 {
			fmt.Println("Cart is empty")
		} else {
			fmt.Printf("---\nSubtotal %s\n---\n", color.Bold(getResponse.Cart.SubTotal.FormattedPrice))
		}
	}
}