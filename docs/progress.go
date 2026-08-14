package docs

var ProgressPage = NewPage(
	"Progress",
	`<p>Progress bars let users track how far a long-running task, such as a file transfer or a multi-step process, has advanced toward completion.</p>`,
	NewExample(
		"Example",
		`<p>Set the <code>value</code> and <code>max</code> attributes to indicate how much of a task has completed out of the total amount of work.</p>`,
		`
    <label>
      Downloading file
      <progress value="70" max="100"></progress>
    </label>
		    `),
	NewExample(
		"Indeterminate",
		`<p>Omit the <code>value</code> attribute to show an indeterminate progress indicator when the completion of a task can't yet be calculated.</p>`,
		`
    <label>
      Downloading file
      <progress></progress>
    </label>
		    `))
