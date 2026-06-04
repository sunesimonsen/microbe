package views

import (
	"fmt"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Spacing() Node {
	sections := []pageSection{}
	for level := range 12 {
		sections = append(sections,
			example(
				fmt.Sprintf("--scale-%d", level),
				Section(
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
