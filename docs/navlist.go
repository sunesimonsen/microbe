package docs

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var NavlistPage = NewPage(
	"Navlist",
	NewExample(
		"Example", `
    <nav class="navlist">
      <details>
        <summary>Content</summary>
        <ul>
          <li><a href="/typography">Typography</a></li>
        </ul>
      </details>
      <details open>
        <summary>Navigation</summary>
        <ul>
          <li><a href="/anchor">Anchor</a></li>
          <li><a href="/navlist" aria-current="page">Navlist</a></li>
        </ul>
      </details>
      <details>
        <summary>Layout</summary>
        <ul>
          <li><a href="/spacing">Spacing</a></li>
        </ul>
      </details>
      <details>
        <summary>Forms</summary>
        <ul>
          <li><a href="/button">Button</a></li>
          <li><a href="/checkbox">Checkbox</a></li>
          <li><a href="/input">Input</a></li>
          <li><a href="/radio">Radio</a></li>
          <li><a href="/range">Range</a></li>
          <li><a href="/select">Select</a></li>
          <li><a href="/switch">Switch</a></li>
        </ul>
      </details>
      <details>
        <summary>Components</summary>
        <ul>
          <li><a href="/card">Card</a></li>
          <li><a href="/dialog">Dialog</a></li>
        </ul>
      </details>
      <details>
        <summary>Theming</summary>
        <ul>
          <li><a href="/accent-color">Accent color</a></li>
          <li><a href="/colors">Colors</a></li>
        </ul>
      </details>
    </nav>
    `,
	),
).WithDescription(
	P(Text("A navlist groups a set of collapsible sections of navigation links, letting users browse and jump between the different areas of an application or site.")),
)
