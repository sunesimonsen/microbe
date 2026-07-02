package docs

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var ProgressPage = NewPage(
	"Progress",
	HGroup(H1(Text("Progress"))),
	NewExample(
		"Progress",
		"Example",
		`
    <label>
      Downloading file
      <progress value="70" max="100"></progress>
    </label>
    `,
	),
	NewExample(
		"Progress",
		"Indeterminate",
		`
    <label>
      Downloading file
      <progress></progress>
    </label>
    `,
	),
)
