package color

import (
	"github.com/codegangsta/cli"
	"github.com/fatih/color"
)

var (
	Bold    = color.New(color.Bold).SprintfFunc()
	Faint   = color.New(color.Faint).SprintfFunc()
	Header  = color.New(color.FgRed, color.Bold).SprintfFunc()
	ShortId = color.New(color.FgBlue).SprintfFunc()
)

func Allow(c *cli.Context) {
	color.NoColor = c.GlobalBool("no-color")
}
