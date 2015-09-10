package color

import (
	"github.com/codegangsta/cli"
	"github.com/fatih/color"
)

var (
	// Bold prints bold text
	Bold = color.New(color.Bold).SprintfFunc()
	// Faint prints faint text
	Faint = color.New(color.Faint).SprintfFunc()
	// Header prints header text
	Header = color.New(color.FgRed, color.Bold).SprintfFunc()
	// ShortID prints short IDs
	ShortID = color.New(color.FgBlue).SprintfFunc()
)

// Allow sets the global output coloring
func Allow(c *cli.Context) {
	color.NoColor = c.GlobalBool("no-color")
}
