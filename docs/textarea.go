package docs

var TextareaPage = NewPage(
	"Textarea",
	`<p>Textareas let users enter and edit multiple lines of free-form text, such as a comment, message, or description, as part of a form.</p>`,
	NewExample(
		"Default",
		`<p>Use a <code>textarea</code> when you need to collect longer, multi-line text such as a comment or description, rather than the single line an <code>input</code> provides.</p>`,
		`
    <label>
      Tell us your story:
      <textarea placeholder="It was a dark and stormy night..."></textarea>
    </label>
    `,
	),
	NewExample(
		"Custom rows",
		`<p>Set the <code>rows</code> attribute to suggest how many lines of text should be visible to the user without needing to scroll.</p>`,
		`
    <label>
      Tell us your story:
      <textarea rows="7" placeholder="It was a dark and stormy night..."></textarea>
    </label>
    `,
	),
	NewExample(
		"Disabled",
		`<p>Add the <code>disabled</code> attribute to prevent a textarea from being edited or focused, and exclude its value from form submission.</p>`,
		`
    <label>
      Tell us your story:
      <textarea disabled placeholder="It was a dark and stormy night..."></textarea>
    </label>
    `,
	),
	NewExample(
		"Read-only",
		`<p>Add the <code>readonly</code> attribute to show text that users can view and select but not edit, while it is still submitted with the form.</p>`,
		`
    <label>
      Tell us your story:
      <textarea readonly placeholder="It was a dark and stormy night...">Read-only value</textarea>
    </label>
    `,
	),
	NewExample(
		"Hint",
		`<p>Place a <code>small</code> element after a textarea to give users extra guidance about what they should write.</p>`,
		`
    <label>
      Tell us your story:
      <textarea placeholder="It was a dark and stormy night..." aria-describedby="scary-story-hint"></textarea>
    </label>
    <small id="scary-story-hint">Scary stories are often more engaging</small>
    `,
	),
	NewExample(
		"Validation",
		`<p>You can indicate the validation state of the textarea using the attribute <code>aria-invilid</code> with the values <code>true</code> and <code>false</code>.</p>`,
		`
    <textarea type="text" name="valid" aria-label="Valid" aria-invalid="false" aria-describedby="valid-hint">Valid</textarea>
    <small id="valid-hint">Looks good!</small>
    <textarea type="text" name="invalid" aria-label="Invalid" aria-invalid="true" aria-describedby="invalid-hint">Invalid</textarea>
    <small id="invalid-hint">Please provide a valid value!</small>
    `,
	),
)
