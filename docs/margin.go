package docs

var MarginPage = NewPage(
	"Margin",
	`
  <p>Microbe adds block margin to all native HTML block elements to make it easy to compose a nicely looking document.</p>
  <p>For custom elements like <code>div</code>, you can add a <code>.content</code> class to get the correct spacing.</p>
  `,
	NewExample(
		"Example",
		`<p>Here you can see how the spacing between blocks flows naturally.`,
		`
    <h3>Document</h3>
    <p>
      Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla
      ullamcorper blandit ultricies. Etiam non suscipit felis. Orci varius
      natoque penatibus et magnis dis parturient montes, nascetur ridiculus
      mus. 
    </p>
    <article class="card content">
      <section>
        Content inside of a card.
      </section>
    </article>
    <div class="content">
      Custom element.
    </div>
    <ul>
      <li>Pellentesque at sodales nisl, in eleifend erat</li>
      <li>Quisque eu maximus leo</li>
      <li>Aenean eget turpis et turpis ullamcorper pretium</li>
    </ul>
    <p>
      Pellentesque at sodales nisl, in eleifend erat. Pellentesque porttitor
      lacinia nibh, a dictum orci condimentum sed. Quisque eu maximus leo.
      Vestibulum elit arcu, molestie eget posuere vel, dapibus pharetra elit.
      Aenean eget turpis et turpis ullamcorper pretium.
    </p>
    `,
	),
)
