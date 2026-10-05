package docs

var GridPage = NewPage(
	"Grid",
	`<p>CSS Grid gives you the primitives to build a responsive grid without a library-specific component. Microbe intentionally does not provide a grid module: the number of columns, gaps, and breakpoints are decisions that should fit your content and visual language.</p>`,
	NewExample(
		"Responsive columns",
		`<p>Start with one column, then add columns at widths where the content still has enough room. Use <code>minmax(0, 1fr)</code> to let each track shrink without allowing long content to force the grid wider than its container.</p>`,
		`
    <style>
      .responsive-grid {
        display: grid;
        grid-template-columns: 1fr;
        gap: var(--scale-5);

        & > * {
          padding: var(--scale-5);
          border: var(--border-width) solid hsl(var(--accent-hue) var(--accent-saturation) var(--lightness-4));
          background: hsl(var(--accent-hue) var(--accent-saturation) var(--lightness-1));
        }
      }

      @media (min-width: 40rem) {
        .responsive-grid {
          grid-template-columns: repeat(2, minmax(0, 1fr));
        }
      }

      @media (min-width: 64rem) {
        .responsive-grid {
          grid-template-columns: repeat(3, minmax(0, 1fr));
        }
      }
    </style>
    <div class="responsive-grid">
      <article>First item</article>
      <article>Second item</article>
      <article>Third item</article>
      <article>Fourth item</article>
      <article>Fifth item</article>
      <article>Sixth item</article>
    </div>
    `,
	),
	NewExample(
		"Placement and spanning",
		`<p>Grid also lets individual items span tracks or occupy a specific area. Change the placement at a breakpoint when the smaller layout needs a different rhythm.</p>`,
		`
    <style>
      .placement-grid {
        display: grid;
        grid-template-columns: repeat(4, minmax(0, 1fr));
        gap: var(--scale-5);

        & > * {
          padding: var(--scale-5);
          border: var(--border-width) solid hsl(var(--accent-hue) var(--accent-saturation) var(--lightness-4));
          background: hsl(var(--accent-hue) var(--accent-saturation) var(--lightness-1));

          &.wide {
            grid-column: span 2;
            border-color: hsl(var(--accent-hue) var(--accent-saturation) var(--lightness-5));
            background: hsl(var(--accent-hue) var(--accent-saturation) var(--lightness-2));
          }
        }
      }

      @media (max-width: 40rem) {
        .placement-grid {
          grid-template-columns: repeat(2, minmax(0, 1fr));
        }

        .placement-grid .wide {
          grid-column: 1 / -1;
        }
      }
    </style>
    <div class="placement-grid">
      <article class="wide">First item</article>
      <article>Second item</article>
      <article>Third item</article>
      <article>Forth item</article>
      <article>Fifth item</article>
      <article class="wide">Sixth item</article>
    </div>
    `,
	),
)
