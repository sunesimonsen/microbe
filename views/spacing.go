package views

import (
	"fmt"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Spacing() Node {
	rows := []Node{}
	for level := 0; level < 12; level++ {
		rows = append(rows, Div(
			Attr("key", fmt.Sprintf("%d", level)),
			Class("spacing-row"),
			Span(Class("spacing-label"), Text(fmt.Sprintf("--scale-%d", level))),
			Span(Class("spacing-box"), StyleAttr(fmt.Sprintf("width: var(--scale-%d)", level))),
		))
	}

	return Group([]Node{
		HGroup(
			ID("content-header"),
			H1(Text("Spacing")),
		),
		Section(
			ID("content"),
			Attr("role", "document"),
			Class("spacing"),
			Group(rows),
		),
	})
}
