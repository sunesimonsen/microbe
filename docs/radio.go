package docs

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
		`<p>Group related radio buttons that share the same <code>name</code> attribute inside a <code>fieldset</code> with a <code>legend</code>, and add <code>disabled</code> to any option that shouldn't currently be selectable.</p>`,
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
		`<p>You can indicate the validation state of the radio group using the attribute <code>aria-invilid</code> with the values <code>true</code> and <code>false</code>.</p>`,
	),
).WithDescription(
	`<p>Radio buttons let users pick exactly one option from a small set of mutually exclusive choices.</p>`,
)
