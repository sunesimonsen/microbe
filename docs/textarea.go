package docs

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

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
		P(Text("Use a "), Code(Text("textarea")), Text(" when you need to collect longer, multi-line text such as a comment or description, rather than the single line an "), Code(Text("input")), Text(" provides.")),
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
		P(Text("Set the "), Code(Text("rows")), Text(" attribute to suggest how many lines of text should be visible to the user without needing to scroll.")),
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
		P(Text("Add the "), Code(Text("disabled")), Text(" attribute to prevent a textarea from being edited or focused, and exclude its value from form submission.")),
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
		P(Text("Add the "), Code(Text("readonly")), Text(" attribute to show text that users can view and select but not edit, while it is still submitted with the form.")),
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
		P(Text("Place a "), Code(Text("small")), Text(" element after a textarea to give users extra guidance about what they should write.")),
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
		P(
			Text("You can indicate the validation state of the textarea using the attribute "),
			Code(Text("aria-invilid")),
			Text(" with the values "),
			Code(Text("true")), Text(" and "), Code(Text("false")),
			Text("."),
		),
	),
).WithDescription(
	P(Text("Textareas let users enter and edit multiple lines of free-form text, such as a comment, message, or description, as part of a form.")),
)
