package docs

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var AboutPage = NewPage(
	"About",
	`
  <p>
    Microbe is a CSS framework with a core stylesheet that elegantly styles most native HTML elements, plus optional modules for common user interface styles.
  </p>
  <p>
    The design system is built around the <a href="https://en.wikipedia.org/wiki/Golden_ratio" target="_blank">golden ratio</a> and the <a href="https://developer.mozilla.org/en-US/docs/Learn_web_development/Core/Styling_basics/Values_and_units#relative_length_units" target="_blank">relative units</a> allowing styles to scale beautifully with the surrounding font size.
  </p>
  <p>
    It ships with opinionated defaults, but you can easily override them using <a href="https://developer.mozilla.org/en-US/docs/Web/CSS/Guides/Cascading_variables/Using_custom_properties" target="_blank">CSS variables</a> and <a href="https://developer.mozilla.org/en-US/docs/Web/CSS/Reference/At-rules/@layer" target="_blank">CSS layers</a>.
  </p>`,
	NewStaticPageSection(
		"Usage",
		H2(Text("Usage")),
		P(Text("Pick a release on the "), A(Href("/docs/releases"), Text("release page")), Text(" and just start adding HTML markup described by the examples.")),
	),
	NewExample(
		"Customizable",
		`<p>Change the accent color or the neutral color or override any CSS styles to tweak the appearance.</p>`,
		`
    <form id="getting-started-demo">
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
    <script>
      let index = 0
      setInterval(() => {
        const form = document.getElementById("getting-started-demo")

        index = (index + 1) % 4;
        form.classList.value = "overrides-" + index;
      }, 3000);
    </script>
    <style>
    @layer overrides {
      #getting-started-demo {
        & * {
          transition: all ease-in-out 1s;
        }

        &.overrides-1,
        &.overrides-2 {
          --accent-hue: 170;
          --accent-saturation: 40%;
          & button,
          & input:not([type=checkbox]) {
            border-radius: var(--scale-6);
          }
        }

        &.overrides-2 {
          --accent-hue: 300;

          button.solid {
            box-shadow: 0 0 var(--scale-3) var(--scale-3) hsl(var(--accent-hue) var(--accent-saturation) var(--lightness-3));
          }
        }

        &.overrides-3,
        &.overrides-4 {
          --accent-hue: 20;

          --neutral-hue: 45;
          --neutral-saturation: 20%;

          & button,
          & input {
            border-radius: 0;
          }
        }
      }
    }
    </style>
    `,
	),
	NewExample(
		"Scalable",
		`<p>All styles are built on the <code>em</code> unit and scales with the surrounding font size.</p>`,
		`
    <div hidden>
      <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" fill="currentColor" class="bi bi-hand-thumbs-up" viewBox="0 0 16 16">
        <path id="thumbs-up-icon" d="M8.864.046C7.908-.193 7.02.53 6.956 1.466c-.072 1.051-.23 2.016-.428 2.59-.125.36-.479 1.013-1.04 1.639-.557.623-1.282 1.178-2.131 1.41C2.685 7.288 2 7.87 2 8.72v4.001c0 .845.682 1.464 1.448 1.545 1.07.114 1.564.415 2.068.723l.048.03c.272.165.578.348.97.484.397.136.861.217 1.466.217h3.5c.937 0 1.599-.477 1.934-1.064a1.86 1.86 0 0 0 .254-.912c0-.152-.023-.312-.077-.464.201-.263.38-.578.488-.901.11-.33.172-.762.004-1.149.069-.13.12-.269.159-.403.077-.27.113-.568.113-.857 0-.288-.036-.585-.113-.856a2 2 0 0 0-.138-.362 1.9 1.9 0 0 0 .234-1.734c-.206-.592-.682-1.1-1.2-1.272-.847-.282-1.803-.276-2.516-.211a10 10 0 0 0-.443.05 9.4 9.4 0 0 0-.062-4.509A1.38 1.38 0 0 0 9.125.111zM11.5 14.721H8c-.51 0-.863-.069-1.14-.164-.281-.097-.506-.228-.776-.393l-.04-.024c-.555-.339-1.198-.731-2.49-.868-.333-.036-.554-.29-.554-.55V8.72c0-.254.226-.543.62-.65 1.095-.3 1.977-.996 2.614-1.708.635-.71 1.064-1.475 1.238-1.978.243-.7.407-1.768.482-2.85.025-.362.36-.594.667-.518l.262.066c.16.04.258.143.288.255a8.34 8.34 0 0 1-.145 4.725.5.5 0 0 0 .595.644l.003-.001.014-.003.058-.014a9 9 0 0 1 1.036-.157c.663-.06 1.457-.054 2.11.164.175.058.45.3.57.65.107.308.087.67-.266 1.022l-.353.353.353.354c.043.043.105.141.154.315.048.167.075.37.075.581 0 .212-.027.414-.075.582-.05.174-.111.272-.154.315l-.353.353.353.354c.047.047.109.177.005.488a2.2 2.2 0 0 1-.505.805l-.353.353.353.354c.006.005.041.05.041.17a.9.9 0 0 1-.121.416c-.165.288-.503.56-1.066.56z"/>
      </svg>
    </div>
    <button style="font-size: 0.8rem" class="outline">
      Scalable
      <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" fill="currentColor" class="bi bi-hand-thumbs-up" viewBox="0 0 16 16">
        <use href="#thumbs-up-icon"></use>
      </svg>
    </button>
    <button class="outline">
      Scalable
      <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" fill="currentColor" class="bi bi-hand-thumbs-up" viewBox="0 0 16 16">
        <use href="#thumbs-up-icon"></use>
      </svg>
    </button>
    <button style="font-size: 1.2rem" class="outline">
      Scalable
      <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" fill="currentColor" class="bi bi-hand-thumbs-up" viewBox="0 0 16 16">
        <use href="#thumbs-up-icon"></use>
      </svg>
    </button>
    `,
	).WithClass("grid"),
	NewExample(
		"Computed colors",
		`<p>As colors are based on hue, saturation and ligthness steps, it is often easy to compute colors that works well with the current accent color.</p>`,
		`
    <div class="grid stretch" style="--hue: 210">
      <button class="solid" style="--accent-hue: var(--hue)">Primary</button>
      <button class="solid" style="--accent-hue: calc(mod(var(--hue) + 180 - 15, 360)); --accent-saturation: 40%;"">Secondary</button>
      <button class="solid" style="--accent-hue: calc(mod(var(--hue) + 180 + 15, 360)); --accent-saturation: 40%;"">Tertiary</button>
    </div>
    <div class="grid stretch" style="--hue: 300">
      <button class="solid" style="--accent-hue: var(--hue)">Primary</button>
      <button class="solid" style="--accent-hue: calc(mod(var(--hue) + 180 - 15, 360)); --accent-saturation: 40%;">Secondary</button>
      <button class="solid" style="--accent-hue: calc(mod(var(--hue) + 180 + 15, 360)); --accent-saturation: 40%;">Tertiary</button>
    </div>
    <div class="grid stretch" style="--hue: 40">
      <button class="solid" style="--accent-hue: var(--hue)">Primary</button>
      <button class="solid" style="--accent-hue: calc(mod(var(--hue) + 180 - 15, 360)); --accent-saturation: 40%;">Secondary</button>
      <button class="solid" style="--accent-hue: calc(mod(var(--hue) + 180 + 15, 360)); --accent-saturation: 40%;">Tertiary</button>
    </div>
    `,
	).WithClass("rows"),
	NewStaticPageSection(
		"Acknowledgement",
		H2(Text("Acknowledgement")),
		P(
			Text("Microbe wasn't built in a vacuum; I was heavily inspired by "),
			ExternalLink("https://picocss.com", "PicoCSS"),
			Text(" both for the design of the documentation site and some of the element styling in the library. The look of the elements are mainly inspired by "),
			ExternalLink("https://garden.zendesk.com", "Zendesk Garden"),
			Text("."),
		),
	),
)
