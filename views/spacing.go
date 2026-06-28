package views

import (
	"fmt"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Spacing() Node {
	sections := []IndexedContent{}
	for level := range 14 {
		sections = append(sections,
			NewExample(
				fmt.Sprintf("--scale-%d", level),
				Nodes(
					Class("spacing"),
					Span(Class("spacing-box"), Style(fmt.Sprintf("width: var(--scale-%d)", level))),
				),
			),
		)
	}

	return docpage(
		HGroup(H1(Text("Spacing"))),
		sections...,
	)
}
