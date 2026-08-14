package docs

var RangePage = NewPage(
	"Range",
	NewExample(
		"Example",
		`
    <label>
      Saturation
      <input type="range" value="70" min="0" max="100" aria-describedby="saturation-hint">
    </label>
    <small id="saturation-hint">Accent color Saturation</small>
		    `,
	).WithDescription(
		`<p>Set the <code>min</code>, <code>max</code>, and <code>value</code> attributes to define the interval of numbers a user can choose from and the starting value.</p>`,
	),
	NewExample(
		"Disabled",
		`
    <label>
      Saturation
      <input type="range" value="70" min="0" max="100" disabled>
    </label>
		    `,
	).WithDescription(
		`<p>Add the <code>disabled</code> attribute to prevent a range input from being changed or focused, and exclude its value from form submission.</p>`,
	),
).WithDescription(
	`<p>Range inputs let users pick a numeric value from within a bounded interval by dragging a handle, which can be quicker than typing an exact number.</p>`,
)
