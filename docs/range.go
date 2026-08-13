package docs

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

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
		P(Text("Set the "), Code(Text("min")), Text(", "), Code(Text("max")), Text(", and "), Code(Text("value")), Text(" attributes to define the interval of numbers a user can choose from and the starting value.")),
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
		P(Text("Add the "), Code(Text("disabled")), Text(" attribute to prevent a range input from being changed or focused, and exclude its value from form submission.")),
	),
).WithDescription(
	P(Text("Range inputs let users pick a numeric value from within a bounded interval by dragging a handle, which can be quicker than typing an exact number.")),
)
