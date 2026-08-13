package docs

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var ProgressPage = NewPage(
	"Progress",
	NewExample(
		"Example",
		`
    <label>
      Downloading file
      <progress value="70" max="100"></progress>
    </label>
    `,
	).WithDescription(
		P(Text("Set the "), Code(Text("value")), Text(" and "), Code(Text("max")), Text(" attributes to indicate how much of a task has completed out of the total amount of work.")),
	),
	NewExample(
		"Indeterminate",
		`
    <label>
      Downloading file
      <progress></progress>
    </label>
    `,
	).WithDescription(
		P(Text("Omit the "), Code(Text("value")), Text(" attribute to show an indeterminate progress indicator when the completion of a task can't yet be calculated.")),
	),
).WithDescription(
	P(Text("Progress bars let users track how far a long-running task, such as a file transfer or a multi-step process, has advanced toward completion.")),
)
