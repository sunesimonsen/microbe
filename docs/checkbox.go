package docs

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var CheckboxPage = NewPage(
	"Checkbox",
	NewExample(
		"Example",
		`
    <fieldset>
      <legend>
        Language preferences:
      </legend>
      <label>
        <input type="checkbox" name="english" 
          aria-describedby="default-language-hint"
          checked disabled>
        English
      </label>
      <small id="default-language-hint">
        You can't disable the default language
      </small>
      <label>
        <input type="checkbox" name="french" checked>
        French
      </label>
      <label>
        <input type="checkbox" name="mandarin">
        Mandarin
      </label>
      <label>
        <input type="checkbox" name="thai">
        Thai
      </label>
      <label>
        <input id="indeterminate-checkbox" type="checkbox" name="quenya">
        Quenya
      </label>
      <label>
        <input type="checkbox" name="dothraki" disabled>
        Dothraki
      </label>
    </fieldset>
    `,
	).WithDescription(
		P(Text("Group related checkboxes inside a "), Code(Text("fieldset")), Text(" with a "), Code(Text("legend")), Text(" describing the group, and add "), Code(Text("disabled")), Text(" to any option that a user shouldn't be able to change, such as a mandatory choice.")),
	),
	NewExample(
		"Hint",
		`
    <label>
      <input type="checkbox" name="newsletter" aria-describedby="newsletter-hint" checked>
      Newsletter
    </label>
    <small id="newsletter-hint">
      We will send you a newsletter every week
    </small>
    `,
	).WithDescription(
		P(Text("Associate a checkbox with helper text using "), Code(Text("aria-describedby")), Text(", giving users additional context about the option before they decide.")),
	),
	NewExample(
		"Validation",
		`
    <label>
      <input type="checkbox" name="valid" aria-invalid="false" aria-describedby="valid-hint">
      Valid
    </label>
    <small id="valid-hint">Looks good!</small>
    <label>
      <input type="checkbox" name="invalid" aria-invalid="true" aria-describedby="invalid-hint" checked>
      Invalid
    </label>
    <small id="invalid-hint">Please provide a valid value!</small>
    `,
	).WithDescription(
		P(
			Text("You can indicate the validation state of the checkbox using the attribute "),
			Code(Text("aria-invilid")),
			Text(" with the values "),
			Code(Text("true")), Text(" and "), Code(Text("false")),
			Text("."),
		),
	),
).WithDescription(
	P(Text("Checkboxes let users select one or more independent options from a set, or toggle a single setting on or off.")),
)
