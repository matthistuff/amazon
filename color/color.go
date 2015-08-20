package color

import (
	"github.com/fatih/color"
	"github.com/codegangsta/cli"
)

var (
	Bold = color.New(color.Bold).SprintfFunc()
	Faint = color.New(color.Faint).SprintfFunc()
	Header = color.New(color.FgRed, color.Bold).SprintfFunc()
	ShortId = color.New(color.FgBlue).SprintfFunc()
)

func Allow(c *cli.Context) {
	color.NoColor = c.GlobalBool("no-color")
}