package docs

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var SelectPage = NewPage(
	"Select",
	NewExample(
		"Single value",
		`
    <label>
      Select a number
      <select aria-label="Select a number">
        <option value="one">One</option>
        <option value="two">Two</option>
        <option value="three">Three</option>
      </select>
    </label>
    `,
	).WithDescription(
		P(Text("Use a "), Code(Text("select")), Text(" without the "), Code(Text("multiple")), Text(" attribute when a user should choose exactly one option from the list.")),
	),
	NewExample(
		"Multiple values",
		`
    <label>
      Select colors
      <select aria-label="Select colors" multiple>
        <option value="blue">Blue</option>
        <option value="green">Green</option>
        <option value="orange">Orange</option>
        <option value="purple">Purple</option>
        <option value="red">Red</option>
        <option value="yellow">Yellow</option>
      </select>
    </label>
    `,
	).WithDescription(
		P(Text("Add the "), Code(Text("multiple")), Text(" attribute to let users pick several options from the list at the same time.")),
	),
	NewExample(
		"Disabled",
		`
    <label>
      Cuisine
      <select aria-label="Select you favorite cuisine..." disabled>
        <option selected disabled value="">Select your favorite cuisine...</option>
        <option>Italian</option>
        <option>Japanese</option>
        <option>Indian</option>
        <option>Thai</option>
        <option>French</option>
      </select>
    </label>
    <label>
      Snacks
      <select name="favorite-snacks" aria-label="Select you favorite snacks..." multiple disabled>
        <option disabled value="">Select your favorite snacks...</option>
        <option>Cheese</option>
        <option>Fruits</option>
        <option>Nuts</option>
        <option>Chocolate</option>
        <option>Crackers</option>
      </select>
    </label>
    `,
	).WithDescription(
		P(Text("Add the "), Code(Text("disabled")), Text(" attribute to prevent a select from being opened or changed, and exclude its value from form submission.")),
	),
	NewExample(
		"Hint",
		`
     <label>
      Cuisine
      <select aria-label="Select you favorite cuisine..." aria-describedby="favorite-cuisine-hint">
        <option selected disabled value="">Select your favorite cuisine...</option>
        <option>Italian</option>
        <option>Japanese</option>
        <option>Indian</option>
        <option>Thai</option>
        <option>French</option>
      </select>
    </label>
    <small id="favorite-cuisine-hint">Select you favorite cuisine</small>
    `,
	).WithDescription(
		P(Text("Associate a select with helper text using "), Code(Text("aria-describedby")), Text(", so users get extra context about what they are choosing.")),
	),
	NewExample(
		"Validation",
		`
    <select name="pizza-topping" aria-label="Select your favorite pizza topping..." aria-invalid="false" aria-describedby="success-hint">
      <option disabled>Select your favorite pizza topping...</option>
      <option selected>Pepperoni</option>
      <option>Mushrooms</option>
      <option>Onions</option>
      <option>Green Peppers</option>
      <option>Olives</option>
    </select>
    <small id="success-hint">Great choice!</small>
    <select name="pizza-topping" aria-label="Select your favorite pizza topping..." aria-invalid="true" aria-describedby="failure-hint">
      <option disabled>Select your favorite pizza topping...</option>
      <option>Pepperoni</option>
      <option>Mushrooms</option>
      <option>Onions</option>
      <option>Green Peppers</option>
      <option>Olives</option>
    </select>
    <small id="failure-hint">Please provide a valid value!</small>
    `,
	).WithDescription(
		P(
			Text("You can indicate the validation state of the select using the attribute "),
			Code(Text("aria-invilid")),
			Text(" with the values "),
			Code(Text("true")), Text(" and "), Code(Text("false")),
			Text("."),
		),
	),
).WithDescription(
	P(Text("Selects let users choose one or more values from a predefined list of options, useful when there are too many choices to show as individual controls.")),
)
