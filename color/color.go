package color

import (
	"github.com/fatih/color"
	"github.com/codegangsta/cli"
)

var (
	Bold = color.New(color.Bold).SprintfFunc()
	Header = color.New(color.FgRed, color.Bold).SprintfFunc()
	ShortId = color.New(color.FgMagenta, color.Bold).SprintfFunc()
)

func Allow(c *cli.Context) {
	color.NoColor = c.GlobalBool("no-color")
}