package actions

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/matthistuff/amazon/api"
	"github.com/matthistuff/amazon/config"
)

func LocalesList(c *cli.Context) {
	for _, locale := range api.GetLocales() {
		fmt.Println(locale)
	}
}

func Locale(c *cli.Context) {
	if len(c.Args()) > 0 {
		LocaleSet(c)
	} else {
		LocaleGet(c)
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
