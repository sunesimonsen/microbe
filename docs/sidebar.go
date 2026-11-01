package docs

var SidebarPage = NewPage(
	"Sidebar",
	`<p>A sidebar for application-style navigation, with a sticky header, a scrollable center section, and a sticky footer. Links and buttons are rendered as sidebar items, and <code>details</code>/<code>summary</code> groups the items into collapsible sections.</p>`,
	NewExample(
		"Example",
		`<p>The <code>header</code> and <code>footer</code> stay in place while the <code>section</code> in between scrolls. The current page is marked with <code>aria-current="page"</code>, which highlights the item with the accent color. The sidebar gets its height from its container, so it works both inside a layout region and as a fixed panel.</p>`,
		`
    <style>
      .sidebar-example > .sidebar {
        block-size: 30.5rem;
        inline-size: 14rem;
      }
    </style>
    <div class="sidebar-example border-layout">
      <nav class="sidebar" data-layout="start">
        <header>
          Sticky at the top
        </header>
        <section>
          <ul>
            <li><a href="/docs/about">About</a></li>
            <li><a href="/docs/modules">Modules</a></li>
          </ul>
          <details open="">
            <summary>Layout</summary>
            <ul>
              <li><a href="/docs/margin">Margin</a></li>
              <li><a href="/docs/spacing">Spacing</a></li>
              <li><a href="/docs/container">Container</a></li>
              <li><a href="/docs/grid">Grid</a></li>
              <li><a href="/docs/border-layout">Border layout</a></li>
              <li><a aria-current="page" href="/docs/sidebar">Sidebar</a></li>
              <li><a href="/docs/accordion">Accordion</a></li>
              <li><a href="/docs/card">Card</a></li>
              <li><a href="/docs/tabs">Tabs</a></li>
            </ul>
          </details>
          <details>
            <summary>Content</summary>
            <ul>
              <li><a href="/docs/typography">Typography</a></li>
              <li><a href="/docs/list">List</a></li>
              <li><a href="/docs/Table">Table</a></li>
            </ul>
          </details>
          <details>
            <summary>Navigation</summary>
            <ul>
              <li><a href="/docs/anchor">Anchor</a></li>
              <li><a href="/docs/navlist">Navlist</a></li>
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
          <details open>
            <summary>Theming</summary>
            <ul>
              <li><a href="/docs/colors">Colors</a></li>
              <li><a href="/docs/palette">Palette</a></li>
            </ul>
          </details>
          <ul>
            <li><a href="#profile">User profile</a></li>
            <li><button>Logout</button></li>
          </ul>
        </section>
        <footer>
          Sticky at the bottom
        </footer>
      </nav>
    </div>
    `,
	).WithModules("Sidebar", "Navlist", "Anchor", "Button", "Border layout"),
)
