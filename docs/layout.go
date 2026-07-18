package docs

import (
	"fmt"
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
	createScaleExamples()...,
)
