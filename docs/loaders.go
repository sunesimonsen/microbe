package docs

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
	),
	NewExample(
		"Indeterminate",
		`
    <label>
      Downloading file
      <progress></progress>
    </label>
    `,
	),
)
