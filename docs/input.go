package docs

var InputPage = NewPage(
	"Input",
	`<p>Inputs let users enter and edit a single line of text or other typed data, such as an email address, number, or password, typically as part of a form.</p>`,
	NewExample(
		"Form",
		`<p>Combine inputs with labels inside a <code>fieldset</code> to build a complete form that users can fill out and then submit or reset.</p>`,
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
	),
	NewExample(
		"Text inputs",
		`<p>Use the <code>text</code>, <code>email</code>, <code>number</code>, <code>password</code>, <code>tel</code> and <code>url</code> types to collect different kinds of single-line text, taking advantage of type-specific validation and on-screen keyboards.</p>`,
		`
    <input type="text" name="text" aria-label="Text" placeholder="Text">
    <input type="email" name="email" aria-label="Email" placeholder="Email" autocomplete="email">
    <input type="number" name="number" aria-label="Number" placeholder="Number">
    <input type="password" name="password" aria-label="Password" placeholder="Password">
    <input type="tel" name="tel" aria-label="Tel" placeholder="Tel">
    <input type="url" name="url" aria-label="Url" placeholder="Url">
    `,
	).WithClass("rows"),
	NewExample(
		"Date and time inputs",
		`<p>Use the <code>date</code>, <code>datetime-local</code>, <code>month</code>, <code>week</code> and <code>time</code> types to let users pick calendar dates and times with the browser's built-in picker instead of typing a specific format.</p>`,
		`
    <input type="date" name="date" aria-label="Date">
    <input type="datetime-local" name="datetime-local" aria-label="Datetime local">
    <input type="month" name="month" aria-label="Month">
    <input type="week" name="week" aria-label="Week">
    <input type="time" name="time" aria-label="Time">
    `,
	).WithClass("rows"),
	NewExample(
		"Search input",
		`<p>Use the <code>search</code> type for a field dedicated to searching content, which some browsers present with a built-in clear button.</p>`,
		`
    <input type="search" name="search" aria-label="Search" placeholder="Search">
    `,
	),
	NewExample(
		"File input",
		`<p>Use the <code>file</code> type to let users choose one or more files from their device to upload, adding the <code>multiple</code> attribute to allow selecting several files at once.</p>`,
		`
    <input type="file" class="solid" tabindex="0" aria-label="File">
    <input type="file" class="solid" tabindex="0" aria-label="File" multiple>
    <input type="file" class="solid" aria-label="File" disabled>
    `,
	).WithClass("rows"),
	NewExample(
		"Color input",
		`<p>Use the <code>color</code> type to let users pick a color using the browser's built-in color picker.</p>`,
		`
    <input type="color" name="color" aria-label="Color" placeholder="Color">
    `,
	),
	NewExample(
		"Disabled",
		`<p>Add the <code>disabled</code> attribute to prevent an input from being edited or focused, and exclude its value from form submission.</p>`,
		`
    <input type="text" name="text" aria-label="Text" placeholder="Text" disabled>
    `,
	),
	NewExample(
		"Read-only",
		`<p>Add the <code>readonly</code> attribute to show a value that users can view and select but not edit, while it is still submitted with the form.</p>`,
		`
    <input type="text" name="text" aria-label="Text" placeholder="Text" readonly value="Read-only value">
    `,
	),
	NewExample(
		"Hint",
		`<p>Place a <code>small</code> element next to an input to give users extra guidance about the expected value before they start typing.</p>`,
		`
    <input type="text" name="Hint" aria-label="Hint" placeholder="Email" aria-describedby="email-hint">
    <small id="email-hint">We’ll never share your email with anyone else.</small>
    `,
	),
	NewExample(
		"Validation",
		`<p>You can indicate the validation state of the input using the attribute <code>aria-invilid</code> with the values <code>true</code> and <code>false</code>.</p>`,
		`
    <input type="text" name="valid" aria-label="Valid" aria-invalid="false" aria-describedby="valid-hint" value="Valid">
    <small id="valid-hint">Looks good!</small>
    <input type="text" name="invalid" aria-label="Invalid" aria-invalid="true" aria-describedby="invalid-hint" value="Invalid">
    <small id="invalid-hint">Please provide a valid value!</small>
    `,
	),
)
