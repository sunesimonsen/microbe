package docs

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var RadioPage = NewPage(
	"Radio",
	NewExample(
		"Example",
		`
    <fieldset>
      <legend>Language preferences:</legend>
      <label>
        <input type="radio" name="language" checked aria-describedby="default-language-hint">
        English
      </label>
      <small id="default-language-hint">Supports all features</small>
      <label>
        <input type="radio" name="language">
        French
      </label>
      <label>
        <input type="radio" name="language">
        Mandarin
      </label>
      <label>
        <input type="radio" name="language">
        Thai
      </label>
      <label>
        <input type="radio" name="language" disabled>
        Dothraki
      </label>
    </fieldset>
    `,
	).WithDescription(
		P(Text("Group related radio buttons that share the same "), Code(Text("name")), Text(" attribute inside a "), Code(Text("fieldset")), Text(" with a "), Code(Text("legend")), Text(", and add "), Code(Text("disabled")), Text(" to any option that shouldn't currently be selectable.")),
	),
	NewExample(
		"Validation",
		`
    <label>
      <input type="radio" name="validation" aria-invalid="false" aria-describedby="valid-hint">
      Valid
    </label>
    <small id="valid-hint">Looks good!</small>
    <label>
      <input type="radio" name="validation" aria-invalid="true" aria-describedby="invalid-hint" checked>
      Invalid
    </label>
    <small id="invalid-hint">This cobination is not allowed!</small>
    `,
	).WithDescription(
		P(
			Text("You can indicate the validation state of the radio group using the attribute "),
			Code(Text("aria-invilid")),
			Text(" with the values "),
			Code(Text("true")), Text(" and "), Code(Text("false")),
			Text("."),
		),
	),
).WithDescription(
	P(Text("Radio buttons let users pick exactly one option from a small set of mutually exclusive choices.")),
)
