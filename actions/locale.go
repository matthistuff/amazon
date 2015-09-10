package actions

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/matthistuff/amazon/api"
	"github.com/matthistuff/amazon/config"
)

// LocalesList lists all available locales
func LocalesList(c *cli.Context) {
	for _, locale := range api.GetLocales() {
		fmt.Println(locale)
	}
}

// Locale sets or prints the current locale
func Locale(c *cli.Context) {
	if len(c.Args()) > 0 {
		LocaleSet(c)
	} else {
		LocaleGet(c)
	}
}

// LocaleSet sets the current locale
func LocaleSet(c *cli.Context) {
	conf := config.GetConfig()
	defer conf.Flush()

	locale := c.Args().First()

	conf.Locale = locale
}

// LocaleGet prints the current locale
func LocaleGet(c *cli.Context) {
	conf := config.GetConfig()
	fmt.Println(conf.Locale)
}
