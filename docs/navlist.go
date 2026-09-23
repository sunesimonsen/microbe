package docs

var NavlistPage = NewPage(
	"Navlist",
	`<p>A navlist styles navigation links and can group them into collapsible sections, letting users browse and jump between the different areas of an application or site.</p>`,
	NewExample(
		"Collapsible sections",
		`<p>Use <code>details</code> and <code>summary</code> to group navigation links into collapsible sections.</p>`, `
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
        <summary>Actions</summary>
        <ul>
	          <li><a href="/docs/button">Button</a></li>
        </ul>
      </details>
      <details>
        <summary>Forms</summary>
        <ul>
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
    `,
	).WithModules("Navlist", "Anchor"),
	NewExample(
		"Simple list",
		`<p>For navigation that does not need collapsible sections, use a regular list inside an element with the <code>navlist</code> class.</p>`, `
    <nav class="navlist">
      <ul>
        <li><a href="/docs/about">About</a></li>
        <li><a href="/docs/modules">Modules</a></li>
        <li><a href="/docs/navlist" aria-current="page">Navlist</a></li>
      </ul>
    </nav>
    `,
	).WithModules("Navlist", "Anchor"),
)
