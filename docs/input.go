package docs

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var InputPage = NewPage(
	"Input",
	NewExample(
		"Form",
		`
    <form>
      <fieldset>
        <label>
          Name
          <input name="name" placeholder="Name" autocomplete="name">
        </label>
        <label>
          Email
          <input name="email" placeholder="Email" autocomplete="email" aria-describedby="email-hint">
          <small id="email-hint">
            We’ll never share your email with anyone else.
          </small>
        </label>
        <label>
          <input type="checkbox" name="newsletter" aria-describedby="newsletter-hint" checked>
          Newsletter
        </label>
        <small id="newsletter-hint">
          We will send you a newsletter every week
        </small>
      </fieldset>
      <div class="actions">
        <button class="outline" type="reset">Reset</button>
        <button class="solid" type="submit">Submit</button>
      </div>
    </form>
    `,
	).WithDescription(
		P(Text("Combine inputs with labels inside a "), Code(Text("fieldset")), Text(" to build a complete form that users can fill out and then submit or reset.")),
	),
	NewExample(
		"Text inputs",
		`
    <input type="text" name="text" aria-label="Text" placeholder="Text">
    <input type="email" name="email" aria-label="Email" placeholder="Email" autocomplete="email">
    <input type="number" name="number" aria-label="Number" placeholder="Number">
    <input type="password" name="password" aria-label="Password" placeholder="Password">
    <input type="tel" name="tel" aria-label="Tel" placeholder="Tel">
    <input type="url" name="url" aria-label="Url" placeholder="Url">
    `,
	).WithDescription(
		P(Text("Use the "), InlineCodeList("text", "email", "number", "password", "tel", "url"), Text(" types to collect different kinds of single-line text, taking advantage of type-specific validation and on-screen keyboards.")),
	).WithClass("rows"),
	NewExample(
		"Date and time inputs",
		`
    <input type="date" name="date" aria-label="Date">
    <input type="datetime-local" name="datetime-local" aria-label="Datetime local">
    <input type="month" name="month" aria-label="Month">
    <input type="week" name="week" aria-label="Week">
    <input type="time" name="time" aria-label="Time">
    `,
	).WithDescription(
		P(Text("Use the "), InlineCodeList("date", "datetime-local", "month", "week", "time"), Text(" types to let users pick calendar dates and times with the browser's built-in picker instead of typing a specific format.")),
	).WithClass("rows"),
	NewExample(
		"Search input",
		`
    <input type="search" name="search" aria-label="Search" placeholder="Search">
    `,
	).WithDescription(
		P(Text("Use the "), Code(Text("search")), Text(" type for a field dedicated to searching content, which some browsers present with a built-in clear button.")),
	),
	NewExample(
		"File input",
		`
    <input type="file" class="solid" tabindex="0" aria-label="File">
    <input type="file" class="solid" tabindex="0" aria-label="File" multiple>
    <input type="file" class="solid" aria-label="File" disabled>
    `,
	).WithDescription(
		P(Text("Use the "), Code(Text("file")), Text(" type to let users choose one or more files from their device to upload, adding the "), Code(Text("multiple")), Text(" attribute to allow selecting several files at once.")),
	).WithClass("rows"),
	NewExample(
		"Color input",
		`
    <input type="color" name="color" aria-label="Color" placeholder="Color">
    `,
	).WithDescription(
		P(Text("Use the "), Code(Text("color")), Text(" type to let users pick a color using the browser's built-in color picker.")),
	),
	NewExample(
		"Disabled",
		`
    <input type="text" name="text" aria-label="Text" placeholder="Text" disabled>
    `,
	).WithDescription(
		P(Text("Add the "), Code(Text("disabled")), Text(" attribute to prevent an input from being edited or focused, and exclude its value from form submission.")),
	),
	NewExample(
		"Read-only",
		`
    <input type="text" name="text" aria-label="Text" placeholder="Text" readonly value="Read-only value">
    `,
	).WithDescription(
		P(Text("Add the "), Code(Text("readonly")), Text(" attribute to show a value that users can view and select but not edit, while it is still submitted with the form.")),
	),
	NewExample(
		"Hint",
		`
    <input type="text" name="Hint" aria-label="Hint" placeholder="Email" aria-describedby="email-hint">
    <small id="email-hint">We’ll never share your email with anyone else.</small>
    `,
	).WithDescription(
		P(Text("Place a "), Code(Text("small")), Text(" element next to an input to give users extra guidance about the expected value before they start typing.")),
	),
	NewExample(
		"Validation",
		`
    <input type="text" name="valid" aria-label="Valid" aria-invalid="false" aria-describedby="valid-hint" value="Valid">
    <small id="valid-hint">Looks good!</small>
    <input type="text" name="invalid" aria-label="Invalid" aria-invalid="true" aria-describedby="invalid-hint" value="Invalid">
    <small id="invalid-hint">Please provide a valid value!</small>
    `,
	).WithDescription(
		P(
			Text("You can indicate the validation state of the input using the attribute "),
			Code(Text("aria-invilid")),
			Text(" with the values "),
			Code(Text("true")), Text(" and "), Code(Text("false")),
			Text("."),
		),
	),
).WithDescription(
	P(Text("Inputs let users enter and edit a single line of text or other typed data, such as an email address, number, or password, typically as part of a form.")),
)
