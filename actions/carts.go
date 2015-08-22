package actions

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/matthistuff/amazon/color"
	"github.com/matthistuff/amazon/config"
	"os"
	"strconv"
)

func CartsDestroy(c *cli.Context) {
	conf := config.GetConfig()
	defer conf.Flush()

	cartName := conf.CartNameFromCache(c.Args().First())

	if cart, exists := conf.Carts[cartName]; exists {
		delete(conf.Carts, cartName)

		fmt.Printf("Deleted cart %s\n", cart.Name)
	} else {
		fmt.Fprintf(os.Stderr, "Cart %s is unknown\n", cartName)
		os.Exit(1)
	}
}

func CartsList(c *cli.Context) {
	color.Allow(c)

	conf := config.GetConfig()
	defer conf.Flush()

	index := 1
	cache := make(map[string]string)
	for _, cart := range conf.Carts {
		fmt.Printf("(%s) %s\n", color.ShortId(strconv.Itoa(index)), cart.Name)
		cache[strconv.Itoa(index)] = cart.Name
		index += 1
	}
	conf.ResultCache["Carts"] = cache
}
