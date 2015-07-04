package actions
import (
	"github.com/codegangsta/cli"
	"github.com/matthistuff/amazon/api"
	"github.com/matthistuff/amazon/config"
	"github.com/pkg/browser"
	"fmt"
	"os"
)

func Checkout(c *cli.Context) {
	cartName := c.Args().First()
	if cartName == "" {
		cartName = "default"
	}

	api := api.Create(c.GlobalString("locale"))

	conf := config.GetConfig()
	defer conf.Flush()

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