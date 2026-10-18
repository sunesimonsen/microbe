package docs

var BorderLayoutPage = NewPage(
	"Border layout",
	`<p>A CSS Grid based layout module in the spirit of the classic <a href="https://docs.oracle.com/javase/8/docs/api/java/awt/BorderLayout.html">BorderLayout</a>. Add the <code>grid-layout</code> class to a container and place children in the <code>top</code>, <code>start</code>, <code>main</code>, <code>end</code>, and <code>bottom</code> areas with the <code>data-layout</code> attribute. Children without a <code>data-layout</code> attribute land in the <code>main</code> area.</p>`,
	NewExample(
		"The five regions",
		`<p>The middle row and column stretch to fill the remaining space, while <code>top</code>, <code>bottom</code>, <code>start</code>, and <code>end</code> are sized by their content. The <code>start</code> and <code>end</code> names follow CSS logical directions, so they adapt if you ever change writing direction.</p>`,
		`
    <style>
      .five-regions {
        gap: var(--scale-4);
        min-block-size: 26rem;

        & > * {
          padding: var(--scale-4);
          border: var(--border-width) solid hsl(var(--accent-hue) var(--accent-saturation) var(--lightness-4));
          background: hsl(var(--accent-hue) var(--accent-saturation) var(--lightness-1));
        }

        & > [data-layout="main"] {
          border-color: hsl(var(--accent-hue) var(--accent-saturation) var(--lightness-5));
          background: hsl(var(--accent-hue) var(--accent-saturation) var(--lightness-2));
        }
      }
    </style>
    <div class="grid-layout five-regions">
      <header data-layout="top">Top</header>
      <aside data-layout="start">Start</aside>
      <main data-layout="main">Main</main>
      <aside data-layout="end">End</aside>
      <footer data-layout="bottom">Bottom</footer>
    </div>
    `,
	).WithModules("Border layout"),
	NewExample(
		"Sparse layouts",
		`<p>Regions you leave empty simply collapse, so the same container works for smaller layouts. Here only <code>top</code>, <code>start</code>, and <code>main</code> are used; the bottom row and end column shrink to nothing.</p>`,
		`
    <style>
      .sparse-regions {
        gap: var(--scale-4);
        min-block-size: 16rem;

        & > * {
          padding: var(--scale-4);
          border: var(--border-width) solid hsl(var(--accent-hue) var(--accent-saturation) var(--lightness-4));
          background: hsl(var(--accent-hue) var(--accent-saturation) var(--lightness-1));
        }

        & > [data-layout="main"] {
          border-color: hsl(var(--accent-hue) var(--accent-saturation) var(--lightness-5));
          background: hsl(var(--accent-hue) var(--accent-saturation) var(--lightness-2));
        }
      }
    </style>
    <div class="grid-layout sparse-regions">
      <header data-layout="top">Top</header>
      <aside data-layout="start">Start</aside>
      <main data-layout="main">Main</main>
    </div>
    `,
	).WithModules("Border layout"),
	NewExample(
		"Nested layouts",
		`<p>Because the layout is driven by a class, containers can be nested freely. The <code>main</code> region below holds a second <code>grid-layout</code> with its own <code>top</code>, <code>start</code>, and <code>main</code> regions, and the <code>end</code> region of the inner layout stacks a pair of cards.</p>`,
		`
    <style>
      .nested-outer,
      .nested-inner {
        gap: var(--scale-3);
      }

      .nested-outer {
        min-block-size: 30rem;

        & > * {
          padding: var(--scale-4);
          border: var(--border-width) solid hsl(var(--accent-hue) var(--accent-saturation) var(--lightness-4));
          background: hsl(var(--accent-hue) var(--accent-saturation) var(--lightness-1));
        }
      }

      .nested-inner {
        min-block-size: 12rem;

        & > * {
          padding: var(--scale-3);
          border: var(--border-width) solid hsl(var(--accent-hue) var(--accent-saturation) var(--lightness-3));
          background: hsl(var(--accent-hue) var(--accent-saturation) var(--lightness-3));
        }
      }
    </style>
    <div class="grid-layout nested-outer">
      <header data-layout="top">Outer top</header>
      <aside data-layout="start">Outer start</aside>
      <div class="grid-layout nested-inner" data-layout="main">
        <header data-layout="top">Inner top</header>
        <aside data-layout="start">Inner start</aside>
        <main data-layout="main">Inner main</main>
        <aside data-layout="end">
          <article>Card one</article>
          <article>Card two</article>
        </aside>
      </div>
      <footer data-layout="bottom">Outer bottom</footer>
    </div>
    `,
	).WithModules("Border layout"),
)
