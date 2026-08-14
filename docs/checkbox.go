package docs

var CheckboxPage = NewPage(
	"Checkbox",
	`<p>Checkboxes let users select one or more independent options from a set, or toggle a single setting on or off.</p>`,
	NewExample(
		"Example",
		`<p>Group related checkboxes inside a <code>fieldset</code> with a <code>legend</code> describing the group, and add <code>disabled</code> to any option that a user shouldn't be able to change, such as a mandatory choice.</p>`,
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
		    `),
	NewExample(
		"Hint",
		`<p>Associate a checkbox with helper text using <code>aria-describedby</code>, giving users additional context about the option before they decide.</p>`,
		`
    <label>
      <input type="checkbox" name="newsletter" aria-describedby="newsletter-hint" checked>
      Newsletter
    </label>
    <small id="newsletter-hint">
      We will send you a newsletter every week
    </small>
		    `),
	NewExample(
		"Validation",
		`<p>You can indicate the validation state of the checkbox using the attribute <code>aria-invilid</code> with the values <code>true</code> and <code>false</code>.</p>`,
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
		    `))
