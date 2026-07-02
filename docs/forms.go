package docs

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var ButtonPage = NewPage(
	"Button",
	HGroup(H1(Text("Button"))),
	NewExample(
		"Button",
		"Styles",
		`
    <button class="solid">Solid Button</button>
    <button class="outline">Outline Button</button>
    <button class="ghost">Ghost Button</button>
    `,
		WithClass("grid"),
		WithDescription(
			P(
				Text("Buttons are not styled by default to avoid the need to reset styles for custom usages. "),
				InlineCodeList(".solid", ".outline", ".ghost"),
				Text(" to set the appearance of the button."),
			),
		),
	),
	NewExample(
		"Button",
		"Disabled",
		`
    <button class="solid" disabled>Solid Button</button>
    <button class="outline" disabled>Outline Button</button>
    <button class="ghost" disabled>Ghost Button</button>
    `,
		WithClass("grid"),
	),
	NewExample(
		"Button",
		"Media",
		`
    <button class="outline">
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-gear" viewBox="0 0 16 16">
        <path d="M8 4.754a3.246 3.246 0 1 0 0 6.492 3.246 3.246 0 0 0 0-6.492M5.754 8a2.246 2.246 0 1 1 4.492 0 2.246 2.246 0 0 1-4.492 0"></path>
        <path d="M9.796 1.343c-.527-1.79-3.065-1.79-3.592 0l-.094.319a.873.873 0 0 1-1.255.52l-.292-.16c-1.64-.892-3.433.902-2.54 2.541l.159.292a.873.873 0 0 1-.52 1.255l-.319.094c-1.79.527-1.79 3.065 0 3.592l.319.094a.873.873 0 0 1 .52 1.255l-.16.292c-.892 1.64.901 3.434 2.541 2.54l.292-.159a.873.873 0 0 1 1.255.52l.094.319c.527 1.79 3.065 1.79 3.592 0l.094-.319a.873.873 0 0 1 1.255-.52l.292.16c1.64.893 3.434-.902 2.54-2.541l-.159-.292a.873.873 0 0 1 .52-1.255l.319-.094c1.79-.527 1.79-3.065 0-3.592l-.319-.094a.873.873 0 0 1-.52-1.255l.16-.292c.893-1.64-.902-3.433-2.541-2.54l-.292.159a.873.873 0 0 1-1.255-.52zm-2.633.283c.246-.835 1.428-.835 1.674 0l.094.319a1.873 1.873 0 0 0 2.693 1.115l.291-.16c.764-.415 1.6.42 1.184 1.185l-.159.292a1.873 1.873 0 0 0 1.116 2.692l.318.094c.835.246.835 1.428 0 1.674l-.319.094a1.873 1.873 0 0 0-1.115 2.693l.16.291c.415.764-.42 1.6-1.185 1.184l-.291-.159a1.873 1.873 0 0 0-2.693 1.116l-.094.318c-.246.835-1.428.835-1.674 0l-.094-.319a1.873 1.873 0 0 0-2.692-1.115l-.292.16c-.764.415-1.6-.42-1.184-1.185l.159-.291A1.873 1.873 0 0 0 1.945 8.93l-.319-.094c-.835-.246-.835-1.428 0-1.674l.319-.094A1.873 1.873 0 0 0 3.06 4.377l-.16-.292c-.415-.764.42-1.6 1.185-1.184l.292.159a1.873 1.873 0 0 0 2.692-1.115z"></path>
      </svg>
      Settings
    </button>
    <button class="solid">
      Settings
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-gear" viewBox="0 0 16 16">
        <path d="M8 4.754a3.246 3.246 0 1 0 0 6.492 3.246 3.246 0 0 0 0-6.492M5.754 8a2.246 2.246 0 1 1 4.492 0 2.246 2.246 0 0 1-4.492 0"></path>
        <path d="M9.796 1.343c-.527-1.79-3.065-1.79-3.592 0l-.094.319a.873.873 0 0 1-1.255.52l-.292-.16c-1.64-.892-3.433.902-2.54 2.541l.159.292a.873.873 0 0 1-.52 1.255l-.319.094c-1.79.527-1.79 3.065 0 3.592l.319.094a.873.873 0 0 1 .52 1.255l-.16.292c-.892 1.64.901 3.434 2.541 2.54l.292-.159a.873.873 0 0 1 1.255.52l.094.319c.527 1.79 3.065 1.79 3.592 0l.094-.319a.873.873 0 0 1 1.255-.52l.292.16c1.64.893 3.434-.902 2.54-2.541l-.159-.292a.873.873 0 0 1 .52-1.255l.319-.094c1.79-.527 1.79-3.065 0-3.592l-.319-.094a.873.873 0 0 1-.52-1.255l.16-.292c.893-1.64-.902-3.433-2.541-2.54l-.292.159a.873.873 0 0 1-1.255-.52zm-2.633.283c.246-.835 1.428-.835 1.674 0l.094.319a1.873 1.873 0 0 0 2.693 1.115l.291-.16c.764-.415 1.6.42 1.184 1.185l-.159.292a1.873 1.873 0 0 0 1.116 2.692l.318.094c.835.246.835 1.428 0 1.674l-.319.094a1.873 1.873 0 0 0-1.115 2.693l.16.291c.415.764-.42 1.6-1.185 1.184l-.291-.159a1.873 1.873 0 0 0-2.693 1.116l-.094.318c-.246.835-1.428.835-1.674 0l-.094-.319a1.873 1.873 0 0 0-2.692-1.115l-.292.16c-.764.415-1.6-.42-1.184-1.185l.159-.291A1.873 1.873 0 0 0 1.945 8.93l-.319-.094c-.835-.246-.835-1.428 0-1.674l.319-.094A1.873 1.873 0 0 0 3.06 4.377l-.16-.292c-.415-.764.42-1.6 1.185-1.184l.292.159a1.873 1.873 0 0 0 2.692-1.115z"></path>
      </svg>
    </button>
    `,
		WithClass("grid"),
		WithDescription(
			P(Text("Buttons uses flexbox layout to neatly position icons.")),
		),
	),
	NewExample(
		"Button",
		"Icon",
		`
    <button class="solid icon" aria-label="Extra options">
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-list" viewBox="0 0 16 16">
        <path fill-rule="evenodd" d="M2.5 12a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H3a.5.5 0 0 1-.5-.5m0-4a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H3a.5.5 0 0 1-.5-.5m0-4a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H3a.5.5 0 0 1-.5-.5"></path>
      </svg>
    </button>
    <button class="outline icon" aria-label="Extra options">
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-list" viewBox="0 0 16 16">
        <path fill-rule="evenodd" d="M2.5 12a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H3a.5.5 0 0 1-.5-.5m0-4a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H3a.5.5 0 0 1-.5-.5m0-4a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H3a.5.5 0 0 1-.5-.5"></path>
      </svg>
    </button>
    <button class="ghost icon" aria-label="Extra options">
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-list" viewBox="0 0 16 16">
        <path fill-rule="evenodd" d="M2.5 12a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H3a.5.5 0 0 1-.5-.5m0-4a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H3a.5.5 0 0 1-.5-.5m0-4a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H3a.5.5 0 0 1-.5-.5"></path>
      </svg>
    </button>
    `,
		WithClass("grid small"),
		WithDescription(
			P(Text("Use the "), Code(Text(".icon")), Text(" for buttons with only an icon.")),
		),
	),
	NewExample(
		"Button",
		"Types",
		`
    <div class="grid stretch">
      <button class="solid">Button</button>
      <button class="outline">Button</button>
      <button class="ghost">Button</button>
    </div>
    <div class="grid stretch">
      <input class="solid" type="button" value="Input input">
      <input class="outline" type="button" value="Input input">
      <input class="ghost" type="button" value="Input input">
    </div>
    <div class="grid stretch">
      <input class="solid" type="submit" value="Summit input">
      <input class="outline" type="submit" value="Summit input">
      <input class="ghost" type="submit" value="Summit input">
    </div>
    <div class="grid stretch">
      <input class="solid" type="reset" value="Reset input">
      <input class="outline" type="reset" value="Reset input">
      <input class="ghost" type="reset" value="Reset input">
    </div>
    `,
		WithClass("rows"),
		WithDescription(
			P(
				Text("In addition to the "), InlineCodeList("button"), Text(" element, input elements with a type of "),
				InlineCodeList("button", "submit", "reset"),
				Text(" can be styled as a button using the classes "),
				InlineCodeList(".solid", ".outline", ".ghost"),
				Text("."),
			),
		),
	),
)

var CheckboxPage = NewPage(
	"Checkbox",
	HGroup(H1(Text("Checkbox"))),
	NewExample(
		"Checkbox",
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
	),
	NewExample(
		"Checkbox",
		"Hint",
		`
    <label>
      <input type="checkbox" name="newsletter" aria-describedby="newsletter-hint" checked>
      Newsletter
    </label>
    <small id="newsletter-hint">
      We will send you a news letter every week
    </small>
    `,
	),
	NewExample(
		"Checkbox",
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
		WithDescription(
			P(
				Text("You can indicate the validation state of the checkbox using the attribute "),
				Code(Text("aria-invilid")),
				Text(" with the values "),
				Code(Text("true")), Text(" and "), Code(Text("false")),
				Text("."),
			),
		),
	),
)

var InputPage = NewPage(
	"Input",
	HGroup(H1(Text("Input"))),
	NewExample(
		"Input",
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
          We will send you a news letter every week
        </small>
      </fieldset>
      <div class="actions">
        <button class="outline" type="reset">Reset</button>
        <button class="solid" type="submit">Submit</button>
      </div>
    </form>
    `,
	),
	NewExample(
		"Input",
		"Text inputs",
		`
    <input type="text" name="text" aria-label="Text" placeholder="Text">
    <input type="email" name="email" aria-label="Email" placeholder="Email" autocomplete="email">
    <input type="number" name="number" aria-label="Number" placeholder="Number">
    <input type="password" name="password" aria-label="Password" placeholder="Password">
    <input type="tel" name="tel" aria-label="Tel" placeholder="Tel">
    <input type="url" name="url" aria-label="Url" placeholder="Url">
    `,
		WithClass("rows"),
	),
	NewExample(
		"Input",
		"Date and time inputs",
		`
    <input type="date" name="date" aria-label="Date">
    <input type="datetime-local" name="datetime-local" aria-label="Datetime local">
    <input type="month" name="month" aria-label="Month">
    <input type="week" name="week" aria-label="Week">
    <input type="time" name="time" aria-label="Time">
    `,
		WithClass("rows"),
	),
	NewExample(
		"Input",
		"Search input",
		`
    <input type="search" name="search" aria-label="Search" placeholder="Search">
    `,
	),
	NewExample(
		"Input",
		"File input",
		`
    <input type="file" class="solid" name="file" aria-label="File" placeholder="File">
    `,
	),
	NewExample(
		"Input",
		"Color input",
		`
    <input type="color" name="color" aria-label="Color" placeholder="Color">
    `,
	),
	NewExample(
		"Input",
		"Disabled",
		`
    <input type="text" name="text" aria-label="Text" placeholder="Text" disabled>
    `,
	),
	NewExample(
		"Input",
		"Read-only",
		`
    <input type="text" name="text" aria-label="Text" placeholder="Text" readonly value="Read-only value">
    `,
	),
	NewExample(
		"Input",
		"Hint",
		`
    <input type="text" name="Hint" aria-label="Hint" placeholder="Email">
    <small>We’ll never share your email with anyone else.</small>
    `,
	),
	NewExample(
		"Input",
		"Validation",
		`
    <input type="text" name="valid" aria-label="Valid" aria-invalid="false" aria-describedby="valid-hint" value="Valid">
    <small id="valid-hint">Looks good!</small>
    <input type="text" name="invalid" aria-label="Invalid" aria-invalid="true" aria-describedby="invalid-hint" value="Invalid">
    <small id="invalid-hint">Please provide a valid value!</small>
    `,
		WithDescription(
			P(
				Text("You can indicate the validation state of the input using the attribute "),
				Code(Text("aria-invilid")),
				Text(" with the values "),
				Code(Text("true")), Text(" and "), Code(Text("false")),
				Text("."),
			),
		),
	),
)

var TextareaPage = NewPage(
	"Textarea",
	HGroup(H1(Text("Textarea"))),
	NewExample(
		"Textarea",
		"Default",
		`
    <label>
      Tell us your story:
      <textarea placeholder="It was a dark and stormy night..."></textarea>
    </label>
    `,
	),
	NewExample(
		"Textarea",
		"Custom rows",
		`
    <label>
      Tell us your story:
      <textarea rows="7" placeholder="It was a dark and stormy night..."></textarea>
    </label>
    `,
	),
	NewExample(
		"Textarea",
		"Disabled",
		`
    <label>
      Tell us your story:
      <textarea disabled placeholder="It was a dark and stormy night..."></textarea>
    </label>
    `,
	),
	NewExample(
		"Textarea",
		"Read-only",
		`
    <label>
      Tell us your story:
      <textarea readonly placeholder="It was a dark and stormy night...">Read-only value</textarea>
    </label>
    `,
	),
	NewExample(
		"Textarea",
		"Hint",
		`
    <label>
      Tell us your story:
      <textarea placeholder="It was a dark and stormy night..."></textarea>
    </label>
    <small>Scary stories are often more engaging</small>
    `,
	),
	NewExample(
		"Textarea",
		"Validation",
		`
    <textarea type="text" name="valid" aria-label="Valid" aria-invalid="false" aria-describedby="valid-hint">Valid</textarea>
    <small id="valid-hint">Looks good!</small>
    <textarea type="text" name="invalid" aria-label="Invalid" aria-invalid="true" aria-describedby="invalid-hint">Invalid</textarea>
    <small id="invalid-hint">Please provide a valid value!</small>
    `,
		WithDescription(
			P(
				Text("You can indicate the validation state of the textarea using the attribute "),
				Code(Text("aria-invilid")),
				Text(" with the values "),
				Code(Text("true")), Text(" and "), Code(Text("false")),
				Text("."),
			),
		),
	),
)

var RadioPage = NewPage(
	"Radio",
	HGroup(H1(Text("Radio"))),
	NewExample(
		"Radio",
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
	),
	NewExample(
		"Radio",
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
		WithDescription(
			P(
				Text("You can indicate the validation state of the radio group using the attribute "),
				Code(Text("aria-invilid")),
				Text(" with the values "),
				Code(Text("true")), Text(" and "), Code(Text("false")),
				Text("."),
			),
		),
	),
)

var RangePage = NewPage(
	"Range",
	HGroup(H1(Text("Range"))),
	NewExample(
		"Range",
		"Example",
		`
    <label>
      Saturation
      <input type="range" value="70" min="0" max="100" aria-describedby="saturation-hint">
    </label>
    <small id="saturation-hint">Accent color Saturation</small>
    `,
	),
	NewExample(
		"Radio",
		"Disabled",
		`
    <label>
      Saturation
      <input type="range" value="70" min="0" max="100" disabled>
    </label>
    `,
	),
)

var SelectPage = NewPage(
	"Select",
	HGroup(H1(Text("Select"))),
	NewExample(
		"Select",
		"Example",
		`
    <label>
      Cuisine
      <select aria-label="Select you favorite cuisine...">
        <option selected disabled value="">
          Select your favorite cuisine...
        </option>
        <option value="0">Italian</option>
        <option value="1">Japanese</option>
        <option value="2">Indian</option>
        <option value="3">Thai</option>
        <option value="4">French</option>
      </select>
    </label>
    <label>
      Snacks
      <select aria-label="Select you favorite snacks..." multiple>
        <option disabled value="">
          Select your favorite snacks...
        </option>
        <option value="cheese">Cheese</option>
        <option value="fruits">Fruits</option>
        <option value="nuts">Nuts</option>
        <option value="chocolate">Chocolate</option>
        <option value="crackers">Crackers</option>
      </select>
    </label>
    `,
	),
	NewExample(
		"Select",
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
	),
	NewExample(
		"Select",
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
	),
	NewExample(
		"Select",
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
	),
	NewExample(
		"Select",
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
	),
	NewExample(
		"Select",
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
	),
)

var SwitchPage = NewPage(
	"Switch",
	HGroup(H1(Text("Switch"))),
	NewExample(
		"Switch",
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
	),
	NewExample(
		"Switch",
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
	),
	NewExample(
		"Switch",
		"Hint",
		`
    <label>
      <input type="checkbox" name="newsletter" role="switch" aria-describedby="newsletter-hint" checked>
      Newsletter
    </label>
    <small id="newsletter-hint">We will send you a news letter every week</small>
    `,
	),
	NewExample(
		"Switch",
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
		WithDescription(
			P(
				Text("You can indicate the validation state of the switch using the attribute "),
				Code(Text("aria-invilid")),
				Text(" with the values "),
				Code(Text("true")), Text(" and "), Code(Text("false")),
				Text("."),
			),
		),
	),
)
