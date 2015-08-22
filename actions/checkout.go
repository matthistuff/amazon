package actions

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/matthistuff/amazon/api"
	"github.com/matthistuff/amazon/config"
	"github.com/pkg/browser"
	"os"
)

func Checkout(c *cli.Context) {
	api := api.Create(c.GlobalString("locale"))

	conf := config.GetConfig()
	defer conf.Flush()

	cartName := conf.CartNameFromCache(c.Args().First())

	if cart, exists := conf.Carts[cartName]; exists {
		if getResponse, getErr := api.CartGet(cart.CartId, cart.HMAC); getErr == nil {
			delete(conf.Carts, cartName)
			browser.OpenURL(getResponse.Cart.PurchaseURL)
		} else {
			panic(getErr)
			return
		}
	} else {
		fmt.Fprintf(os.Stderr, "Cart %s is unknown\n", cartName)
		os.Exit(1)
	}
}
