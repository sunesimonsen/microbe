package docs

import (
	"fmt"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func createScaleExamples() []PageSection {
	scaleExamples := []PageSection{}

	for level := range 14 {
		scaleExamples = append(scaleExamples,
			NewExample(
				fmt.Sprintf("--scale-%d", level),
				fmt.Sprintf("<span class=\"spacing-box\" style=\"width: var(--scale-%d)\"></span>", level),
			),
		)
	}

	return scaleExamples
}

var SpacingPage = NewPage(
	"Spacing",
	HGroup(H1(Text("Spacing"))),
	createScaleExamples()...,
)
