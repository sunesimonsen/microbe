package docs

var TextareaPage = NewPage(
	"Textarea",
	NewExample(
		"Default",
		`
    <label>
      Tell us your story:
      <textarea placeholder="It was a dark and stormy night..."></textarea>
    </label>
		    `,
	).WithDescription(
		`<p>Use a <code>textarea</code> when you need to collect longer, multi-line text such as a comment or description, rather than the single line an <code>input</code> provides.</p>`,
	),
	NewExample(
		"Custom rows",
		`
    <label>
      Tell us your story:
      <textarea rows="7" placeholder="It was a dark and stormy night..."></textarea>
    </label>
		    `,
	).WithDescription(
		`<p>Set the <code>rows</code> attribute to suggest how many lines of text should be visible to the user without needing to scroll.</p>`,
	),
	NewExample(
		"Disabled",
		`
    <label>
      Tell us your story:
      <textarea disabled placeholder="It was a dark and stormy night..."></textarea>
    </label>
		    `,
	).WithDescription(
		`<p>Add the <code>disabled</code> attribute to prevent a textarea from being edited or focused, and exclude its value from form submission.</p>`,
	),
	NewExample(
		"Read-only",
		`
    <label>
      Tell us your story:
      <textarea readonly placeholder="It was a dark and stormy night...">Read-only value</textarea>
    </label>
		    `,
	).WithDescription(
		`<p>Add the <code>readonly</code> attribute to show text that users can view and select but not edit, while it is still submitted with the form.</p>`,
	),
	NewExample(
		"Hint",
		`
    <label>
      Tell us your story:
      <textarea placeholder="It was a dark and stormy night..." aria-describedby="scary-story-hint"></textarea>
    </label>
    <small id="scary-story-hint">Scary stories are often more engaging</small>
		    `,
	).WithDescription(
		`<p>Place a <code>small</code> element after a textarea to give users extra guidance about what they should write.</p>`,
	),
	NewExample(
		"Validation",
		`
    <textarea type="text" name="valid" aria-label="Valid" aria-invalid="false" aria-describedby="valid-hint">Valid</textarea>
    <small id="valid-hint">Looks good!</small>
    <textarea type="text" name="invalid" aria-label="Invalid" aria-invalid="true" aria-describedby="invalid-hint">Invalid</textarea>
    <small id="invalid-hint">Please provide a valid value!</small>
		    `,
	).WithDescription(
		`<p>You can indicate the validation state of the textarea using the attribute <code>aria-invilid</code> with the values <code>true</code> and <code>false</code>.</p>`,
	),
).WithDescription(
	`<p>Textareas let users enter and edit multiple lines of free-form text, such as a comment, message, or description, as part of a form.</p>`,
)
