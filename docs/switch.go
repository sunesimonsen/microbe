package docs

var SwitchPage = NewPage(
	"Switch",
	`<p>Switches let users toggle a single setting on or off immediately, without needing to submit a form to see the state change.</p>`,
	NewExample(
		"Example",
		`<p>Add <code>role="switch"</code> to a checkbox to indicate it represents an on/off setting that takes effect immediately, rather than a form field awaiting submission.</p>`,
		`
    <label>
      <input type="checkbox" name="terms" role="switch">
      I agree to the Terms
    </label>
    <label>
      <input type="checkbox" name="opt-in" role="switch" checked>
      Receive news and offers
    </label>
		    `),
	NewExample(
		"Disabled",
		`<p>Add the <code>disabled</code> attribute to prevent a switch from being toggled or focused, and exclude its value from form submission.</p>`,
		`
    <label>
      <input type="checkbox" role="switch" disabled checked>
      Disabled
    </label>
    <label>
      <input type="checkbox" role="switch" disabled>
      Disabled
    </label>
		    `),
	NewExample(
		"Hint",
		`<p>Associate a switch with helper text using <code>aria-describedby</code>, so users understand what the setting controls before toggling it.</p>`,
		`
    <label>
      <input type="checkbox" name="newsletter" role="switch" aria-describedby="newsletter-hint" checked>
      Newsletter
    </label>
    <small id="newsletter-hint">We will send you a newsletter every week</small>
		    `),
	NewExample(
		"Validation",
		`<p>You can indicate the validation state of the switch using the attribute <code>aria-invilid</code> with the values <code>true</code> and <code>false</code>.</p>`,
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
		    `))
