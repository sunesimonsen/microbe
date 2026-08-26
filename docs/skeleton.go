package docs

var SkeletonPage = NewPage(
	"Skeleton",
	`
  <p>A skeleton loader shows users a blank version of a page or section of a page into which content is gradually loaded. It provides a visual estimate of the space needed.</p>
  <p>Skeleton loaders are the same color as the text color at 7% opacity.</p>
  `,
	NewExample(
		"Example",
		``,
		`
    <section id="example-loading">
      <h3><span class="skeleton" style="width: 40%"></span></h3>
      <p><span class="skeleton" style="width: 70%"></span></p>
      <p><span class="skeleton"></span></p>
    </section>
    <section id="example-loaded" hidden>
      <h3>Skeleton loaders</h3>
      <p>The skeleton should mirror the actual design as closely as possible.</p>
      <p>But even if it isn't really precise, it still gives a better experience than spinners in most cases.</p>
    </section>
    <script>
    setTimeout(() => {
      const skeleton = document.getElementById('example-loading');
      const loadedContent = document.getElementById('example-loaded');
      loadedContent.removeAttribute('hidden');
      skeleton.setAttribute('hidden', '')
    }, 3000);
    </script>
    `,
	),
	NewExample(
		"Full width",
		``,
		`
    <h3><span class="skeleton"></span></h3>
    <p><span class="skeleton"></span></p>
    <p><span class="skeleton"></span></p>
    <p><span class="skeleton"></span></p>
    `,
	),
)
