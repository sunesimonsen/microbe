package docs

var ContainerPage = NewPage(
	"Container",
	`<p>A container controls the readable width of a page while allowing the viewport to grow around it. You can build one with CSS Grid and media queries, without adopting a framework-wide container width or breakpoint scale. Microbe intentionally does not provide a container module because those values are part of your page's layout.</p>`,
	NewExample(
		"Centered content",
		`<p>For a simple centered container, let the block fill smaller viewports, cap it with <code>max-width</code>, and center it with <code>margin-inline: auto</code>. Choose breakpoint-specific widths based on when your content needs more room, not on a particular device.</p>`,
		`
    <style>
      .container {
        max-width: 20rem;
        margin-inline: auto;
      }

      :is(.container, .container-fluid) > * {
        padding: var(--scale-5);
        background: hsl(var(--accent-hue) var(--accent-saturation) var(--lightness-1));
        border: var(--border-width) solid hsl(var(--accent-hue) var(--accent-saturation) var(--lightness-4));
        margin-block: var(--scale-5);
      }

      @media (min-width: 80rem) {
        .container {
          max-width: 30rem;
        }
      }

      @media (min-width: 100rem) {
        .container {
          max-width: 40rem;
        }
      }
    </style>
    <div class="container">
      <section>
        <p>The content stays centered and gains space only when the viewport is wide enough.</p>
      </section>
    </div>
    <div class="container-fluid">
      <section>
        <p>The content spans the full width.</p>
      </section>
    </div>
    <div class="container">
      <section>
        <p>The content stays centered and gains space only when the viewport is wide enough.</p>
      </section>
    </div>
    `,
	),
)
