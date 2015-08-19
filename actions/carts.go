package actions

import (
	"github.com/codegangsta/cli"
	"fmt"
	"github.com/matthistuff/amazon/config"
	"strconv"
	"os"
)

func CartsDestroy(c *cli.Context) {
	cartName := c.Args().First()

	conf := config.GetConfig()
	defer conf.Flush()

	if cart, exists := conf.Carts[cartName]; exists {
		delete(conf.Carts, cartName)

		fmt.Printf("Deleted cart %s\n", cart.Name)
	} else {
		fmt.Fprintf(os.Stderr, "Cart %s is unknown\n", cartName)
		os.Exit(1)
	}
}

func CartsList(c *cli.Context) {
	conf := config.GetConfig()
	defer conf.Flush()

	index := 1
	cache := make(map[string]string)
	for _, cart := range conf.Carts {
		fmt.Printf("(%d) %s\n", index, cart.Name)
		cache[strconv.Itoa(index)] = cart.Name
		index +=1
	}
	conf.ResultCache["Carts"] = cache
}