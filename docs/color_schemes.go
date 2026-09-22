package docs

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var ColorSchemesPage = NewPage(
	"Color schemes",
	`<p>Microbe includes light and dark color schemes in the core stylesheet. Schemes change the color tokens used by the framework, including the lightness scale, foreground and background colors, focus colors, and native form control appearance.</p>
  <p>A color scheme is selected by the user's operating system by default, or explicitly with a class. The classes can be applied to the document or to an individual part of an interface, so a page can use the system preference while a component uses a fixed scheme.</p>`,
	NewStaticPageSection(
		"Modes",
		H2(Text("Modes")),
		P(Text("There are three ways to select a scheme:")),
		Ul(
			Li(
				Strong(Text("System (default). ")),
				Text("Do not add a color-scheme class. Microbe starts in light mode and switches to dark mode when the browser matches the prefers-color-scheme: dark media feature."),
			),
			Li(
				Strong(Text("Light. ")),
				Code(Text("color-scheme-light")),
				Text(" forces the light palette and light native controls, regardless of the system preference."),
			),
			Li(
				Strong(Text("Dark. ")),
				Code(Text("color-scheme-dark")),
				Text(" forces the dark palette and dark native controls, regardless of the system preference."),
			),
		),
		P(
			Text("The classes are ordinary CSS classes; Microbe does not require JavaScript to select a scheme. To change a scheme at runtime, add or remove the class with your own application code. Removing both classes returns the element to system mode."),
		),
	),
	NewStaticPageSection(
		"Applying a scheme",
		H2(Text("Applying a scheme")),
		P(Text("Apply a class to the element that owns the interface. Put the class on the document root to control the whole page:")),
		Pre(
			Code(
				Class("language-html"),
				Data("highlight", "yes"),
				Text(`<html class="microbe color-scheme-dark">...</html>`),
			),
		),
		P(Text("A scheme can also be scoped to a part of a page. Descendants inherit its color variables and native control preference:")),
		Pre(
			Code(
				Class("language-html"),
				Data("highlight", "yes"),
				Text(`<section class="color-scheme-light">...</section>`),
			),
		),
		P(
			Text("When using the core stylesheet, make sure the interface is inside a "),
			Code(Text("microbe")),
			Text(" container. This is normally the document root, as in the first example."),
		),
	),
	NewExample(
		"Switching schemes",
		`<p>This example changes the scheme by toggling the classes on a containing element. The selection control is only demo code; an application can use a saved preference, a user setting, or the system preference instead.</p>`,
		`
    <article id="color-scheme-example" class="card color-scheme-dark">
      <header>
        <label>
          Color scheme
          <select id="color-scheme-select" aria-label="Color scheme">
            <option value="system">System</option>
            <option value="light">Light</option>
            <option value="dark" selected>Dark</option>
          </select>
        </label>
      </header>
      <section>
        <form>
          <fieldset>
            <label>
              Name
              <input name="name" placeholder="Name" autocomplete="name">
            </label>
            <label>
              Email
              <input name="email" placeholder="Email" autocomplete="email" aria-describedby="dark-email-hint">
              <small id="dark-email-hint">
                We’ll never share your email with anyone else.
              </small>
            </label>
            <label>
              <input type="checkbox" name="newsletter" aria-describedby="dark-newsletter-hint" checked>
              Newsletter
            </label>
            <small id="dark-newsletter-hint">
              We will send you a newsletter every week
            </small>
          </fieldset>
          <div class="actions">
            <button class="outline" type="reset">Reset</button>
            <button class="solid" type="submit">Submit</button>
          </div>
        </form>
      </section>
    </article>
    <script>
      const colorSchemeExample = document.getElementById("color-scheme-example")
      const colorSchemeSelect = document.getElementById("color-scheme-select")

      colorSchemeSelect.addEventListener("change", (event) => {
        const colorScheme = event.target.value
        colorSchemeExample.classList.toggle("color-scheme-light", colorScheme === "light")
        colorSchemeExample.classList.toggle("color-scheme-dark", colorScheme === "dark")
      })
    </script>
    `,
	).WithModules("Button", "Card", "Forms"),
	NewStaticPageSection(
		"Color tokens",
		H2(Text("Color tokens")),
		P(
			Text("Both built-in schemes use the same token names, so components do not need separate light and dark styles. The scheme changes the values behind these tokens:"),
		),
		Ul(
			Li(Code(Text("--lightness-0")), Text(" through "), Code(Text("--lightness-12")), Text(" define the palette's lightness steps. See the "), A(Href("/docs/palette"), Text("Palette")), Text(" page for the complete scale.")),
			Li(Code(Text("--foreground-lightness")), Text(" and "), Code(Text("--background-lightness")), Text(" define the default text and surface contrast.")),
			Li(Code(Text("--focus-lightness")), Text(", "), Code(Text("--focus-outline-width")), Text(", and "), Code(Text("--focus-outline-offset")), Text(" adapt focus indicators to the selected contrast.")),
			Li(Code(Text("--accent-saturation")), Text(" and "), Code(Text("--error-saturation")), Text(" adjust semantic accent and error colors for the background.")),
			Li(Code(Text("--neutral-hue")), Text(", "), Code(Text("--neutral-saturation")), Text(", and "), Code(Text("--shadow-color")), Text(" adjust neutral surfaces and shadows.")),
		),
		P(Text("The light scheme uses a white background and a black foreground. The dark scheme uses a dark background, a light foreground, a reduced accent saturation, and a neutral hue derived from the accent hue. All Microbe components consume these variables, so the same semantic HTML works in either scheme.")),
	),
	NewStaticPageSection(
		"Browser preference",
		H2(Text("Browser preference")),
		P(
			Text("Microbe sets the CSS color-scheme property for each mode so the browser can style native controls, scrollbars, and other user-agent UI consistently. The system mode follows prefers-color-scheme: dark; an explicit color-scheme-light or color-scheme-dark class takes precedence over that media query."),
		),
		P(Text("Microbe does not store a user's choice or provide a preference switch. If an application offers a selector, it should manage the class and persist the choice as appropriate. The example above demonstrates the smallest possible class-switching implementation.")),
	),
)
