package actions

import (
	"github.com/codegangsta/cli"
	"github.com/matthistuff/amazon/api"
	"fmt"
	"github.com/matthistuff/amazon/config"
)

func LocaleList(c *cli.Context) {
	for _, locale := range api.GetLocales() {
		fmt.Println(locale)
	}
}

func LocaleSet(c *cli.Context) {
	conf := config.GetConfig()
	defer conf.Flush()

	locale := c.Args().First()

	conf.Locale = locale
}

func LocaleGet(c *cli.Context) {
	conf := config.GetConfig()
	fmt.Println(conf.Locale)
}