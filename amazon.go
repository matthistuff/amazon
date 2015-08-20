package main

import (
	"os"
	"github.com/codegangsta/cli"
	"github.com/matthistuff/amazon/actions"
	"github.com/matthistuff/amazon/config"
)

func init() {
	if err := config.LoadConfig(); err != nil {
		panic(err)
	}
}

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = "amazon"
	app.Usage = "CLI interface to amazon.*"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "locale, l",
			Value: "",
			Usage: "Amazon locale",
			EnvVar: "AMAZON_LOCALE",
		},
		cli.BoolFlag{
			Name: "no-color",
			Usage: "disable colored output",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "search",
			Usage: "search for products",
			Action: actions.Search,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name: "page, p",
					Value: 1,
					Usage: "search results page",
				},
			},
		},
		{
			Name: "info",
			Usage: "get product info",
			Action: actions.Info,
		},
		{
			Name: "open",
			Usage: "open product in browser",
			Action: actions.Open,
		},
		{
			Name: "cart",
			Usage: "manage a cart",
			Subcommands: []cli.Command{
				{
					Name: "add",
					Usage: "add item to cart",
					Action: actions.CartAdd,
				},
				{
					Name: "update",
					Usage: "update item in cart",
					Action: actions.CartUpdate,
				},
				{
					Name: "remove",
					Usage: "remove item from cart",
					Action: actions.CartRemove,
				},
				{
					Name: "info",
					Usage: "list cart items",
					Action: actions.CartInfo,
				},
			},
		},
		{
			Name: "carts",
			Usage: "manage carts",
			Subcommands: []cli.Command{
				{
					Name: "list",
					Usage: "list all active varts",
					Action: actions.CartsList,
				},
				{
					Name: "destroy",
					Usage: "delete a cart",
					Action: actions.CartsDestroy,
				},
			},
		},
		{
			Name: "checkout",
			Usage: "proceed to checkout",
			Action: actions.Checkout,
		},
		{
			Name: "locale",
			Usage: "manage locale",
			Subcommands: []cli.Command{
				{
					Name: "list",
					Usage: "list all available locales",
					Action: actions.LocaleList,
				},
				{
					Name: "set",
					Usage: "set locale",
					Action: actions.LocaleSet,
				},
				{
					Name: "get",
					Usage: "get current locale",
					Action: actions.LocaleGet,
				},
			},
		},
	}

	app.Run(os.Args)
}
