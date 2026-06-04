package views

import (
	"fmt"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Spacing() Node {
	rows := []Node{}
	for level := range 12 {
		rows = append(rows,
			example(
				fmt.Sprintf("--scale-%d", level),
				Span(Class("spacing-box"), Style(fmt.Sprintf("width: var(--scale-%d)", level))),
			),
		)
	}

	return Group([]Node{
		HGroup(
			H1(Text("Spacing")),
		),
		Section(
			Role("document"),
			Class("spacing"),
			Group(rows),
		),
	})
}
