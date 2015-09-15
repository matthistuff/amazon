package helper

import (
	"fmt"
	"github.com/matthistuff/amazon/color"
	"github.com/matthistuff/html2text"
	"golang.org/x/net/html"
	"strings"
)

type emFormatter struct{}

func (emFormatter) Format(node *html.Node, childIndex int) (string, error) {
	return color.Bold("%%s"), nil
}

func init() {
	html2text.Formatters["em"] = emFormatter{}
	html2text.Formatters["b"] = emFormatter{}
}

// FormatText prints pretty text
func FormatText(str string, cap int, prepend string) string {
	text, _ := html2text.FromString(str)
	lines := strings.Split(text, "\n")

	for i, line := range lines {
		lines[i] = fmt.Sprintf(
			"%s%s", prepend, strings.Join(
				slashText(line, cap),
				fmt.Sprintf("\n%s", prepend)))
	}

	return strings.Join(lines, "\n")
}

func slashText(str string, cap int) []string {
	var out []string
	letters := []rune(str)

	tmp := ""
	var char string
	var at int
	for i, r := range letters {
		char = string(r)
		tmp = tmp + char

		at = i - len(out)*cap
		if (at >= cap) && (char == " ") {
			out = append(out, tmp)
			tmp = ""
		}
	}

	if len(tmp) > 0 {
		out = append(out, tmp)
	}

	return out
}
