package docs

var NavlistPage = NewPage(
	"Navlist",
	`<p>A navlist groups a set of collapsible sections of navigation links, letting users browse and jump between the different areas of an application or site.</p>`,
	NewExample(
		"Example", "", `
    <nav class="navlist">
      <details>
        <summary>Content</summary>
        <ul>
	          <li><a href="/docs/typography">Typography</a></li>
	          <li><a href="/docs/list">List</a></li>
	          <li><a href="/docs/Table">Table</a></li>
        </ul>
      </details>
      <details open>
        <summary>Navigation</summary>
        <ul>
	          <li><a href="/docs/anchor">Anchor</a></li>
	          <li><a href="/docs/navlist" aria-current="page">Navlist</a></li>
        </ul>
      </details>
      <details>
        <summary>Forms</summary>
        <ul>
	          <li><a href="/docs/button">Button</a></li>
	          <li><a href="/docs/checkbox">Checkbox</a></li>
	          <li><a href="/docs/input">Input</a></li>
	          <li><a href="/docs/radio">Radio</a></li>
	          <li><a href="/docs/range">Range</a></li>
	          <li><a href="/docs/select">Select</a></li>
	          <li><a href="/docs/switch">Switch</a></li>
        </ul>
      </details>
      <details>
        <summary>Theming</summary>
        <ul>
	          <li><a href="/docs/colors">Colors</a></li>
	          <li><a href="/docs/palette">Palette</a></li>
        </ul>
      </details>
    </nav>
    `))
