package docs

var SidebarPage = NewPage(
	"Sidebar",
	`<p>A classic sidebar layout built with the <a href="/docs/border-layout">Border layout</a> module. Place navigation or auxiliary content in the <code>start</code> or <code>end</code> region and let the <code>main</code> region take the remaining space.</p>
<p>Microbe does not ship styling for sidebars out of the box. This kind of styling should adhere to your personal brand, and as the example below shows, it is easy to create your own styles.</p>`,
	NewExample(
		"Example",
		`<p>The sidebar holds the navigation in the <code>start</code> region, while the content in <code>main</code> takes the remaining space. Below a <code>40rem</code> viewport, the grid areas are redefined with a media query, so the sidebar stacks on top of the content instead. The sidebar is styled with the neutral hue, so it adapts to light and dark color schemes.</p>`,
		`
    <style>
      .sidebar-layout {
        gap: var(--scale-4);

        & > aside {
          padding: var(--scale-5);
          background: hsl(var(--neutral-hue) var(--neutral-saturation) var(--lightness-1));
          border: var(--border-width) solid hsl(var(--neutral-hue) var(--neutral-saturation) var(--lightness-3));
          border-radius: var(--scale-3);
          min-width: var(--scale-10);
        }

        & > main {
          padding-inline: var(--scale-5);
        }
      }

      @media (max-width: 40rem) {
        .sidebar-layout {
          grid-template-columns: auto;
          grid-template-rows: auto auto;
          grid-template-areas:
            "start"
            "main";
        }
      }
    </style>
    <div class="grid-layout sidebar-layout">
      <aside data-layout="start">
        <nav>
          <ul>
            <li><a href="#" aria-current="true">Overview</a></li>
            <li><a href="#">Details</a></li>
            <li><a href="#">Settings</a></li>
          </ul>
        </nav>
      </aside>
      <main data-layout="main">
        <h2>Overview</h2>
        <p>Narrow the viewport below 40rem to see the sidebar stack on top of the content.</p>
      </main>
    </div>
    `,
	).WithModules("Border layout"),
)
