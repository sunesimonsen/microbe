package docs

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var SwitchPage = NewPage(
	"Switch",
	NewExample(
		"Example",
		`
    <label>
      <input type="checkbox" name="terms" role="switch">
      I agree to the Terms
    </label>
    <label>
      <input type="checkbox" name="opt-in" role="switch" checked>
      Receive news and offers
    </label>
    `,
	).WithDescription(
		P(Text("Add "), Code(Text("role=\"switch\"")), Text(" to a checkbox to indicate it represents an on/off setting that takes effect immediately, rather than a form field awaiting submission.")),
	),
	NewExample(
		"Disabled",
		`
    <label>
      <input type="checkbox" role="switch" disabled checked>
      Disabled
    </label>
    <label>
      <input type="checkbox" role="switch" disabled>
      Disabled
    </label>
    `,
	).WithDescription(
		P(Text("Add the "), Code(Text("disabled")), Text(" attribute to prevent a switch from being toggled or focused, and exclude its value from form submission.")),
	),
	NewExample(
		"Hint",
		`
    <label>
      <input type="checkbox" name="newsletter" role="switch" aria-describedby="newsletter-hint" checked>
      Newsletter
    </label>
    <small id="newsletter-hint">We will send you a newsletter every week</small>
    `,
	).WithDescription(
		P(Text("Associate a switch with helper text using "), Code(Text("aria-describedby")), Text(", so users understand what the setting controls before toggling it.")),
	),
	NewExample(
		"Validation",
		`
    <label>
      <input type="checkbox" role="switch" name="valid" aria-invalid="false" aria-describedby="valid-hint">
      Valid
    </label>
    <small id="valid-hint">Looks good!</small>
    <label>
      <input type="checkbox" role="switch" name="invalid" aria-invalid="true" aria-describedby="invalid-hint" checked>
      Invalid
    </label>
    <small id="invalid-hint">Please provide a valid value!</small>
    `,
	).WithDescription(
		P(
			Text("You can indicate the validation state of the switch using the attribute "),
			Code(Text("aria-invilid")),
			Text(" with the values "),
			Code(Text("true")), Text(" and "), Code(Text("false")),
			Text("."),
		),
	),
).WithDescription(
	P(Text("Switches let users toggle a single setting on or off immediately, without needing to submit a form to see the state change.")),
)
