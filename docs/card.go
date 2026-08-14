package docs

var CardPage = NewPage(
	"Card",
	NewExample(
		"Bare card",
		`
    <article class="card">
      <section>
        <p>
          Vestibulum venenatis neque nec iaculis viverra. Proin a odio ex.
          Etiam eget rutrum nulla, id finibus tellus. Proin dignissim tortor
          lacus, in vehicula arcu commodo ut. Aenean quis euismod lectus, vitae
          ornare ex. Vivamus posuere lectus at quam maximus semper. Fusce
          sagittis in est a consequat.
        </p>
      </section>
    </article>
		    `,
	).WithDescription(
		`<p>A minimal card that simply groups a single piece of content, such as a short summary or teaser, without any additional header or footer.</p>`,
	),
	NewExample(
		"Card with header",
		`
    <article class="card">
      <header>Header</header>
      <section>
        <p>
          Vestibulum venenatis neque nec iaculis viverra. Proin a odio ex.
          Etiam eget rutrum nulla, id finibus tellus. Proin dignissim tortor
          lacus, in vehicula arcu commodo ut. Aenean quis euismod lectus, vitae
          ornare ex. Vivamus posuere lectus at quam maximus semper. Fusce
          sagittis in est a consequat.
        </p>
      </section>
    </article>
		    `,
	).WithDescription(
		`<p>Add a <code>header</code> to a card to give the grouped content a title, so users can quickly identify what the card is about.</p>`,
	),
	NewExample(
		"Card with footer",
		`
    <article class="card">
      <section>
        <p>
          Vestibulum venenatis neque nec iaculis viverra. Proin a odio ex.
          Etiam eget rutrum nulla, id finibus tellus. Proin dignissim tortor
          lacus, in vehicula arcu commodo ut. Aenean quis euismod lectus, vitae
          ornare ex. Vivamus posuere lectus at quam maximus semper. Fusce
          sagittis in est a consequat.
        </p>
      </section>
      <footer>Footer</footer>
    </article>
		    `,
	).WithDescription(
		`<p>Add a <code>footer</code> to a card to hold supplementary information or actions, such as a call-to-action link, that relate to the card's content.</p>`,
	),
	NewExample(
		"Card with header and footer",
		`
    <article class="card">
      <header>Header</header>
      <section>
        <p>
          Vestibulum venenatis neque nec iaculis viverra. Proin a odio ex.
          Etiam eget rutrum nulla, id finibus tellus. Proin dignissim tortor
          lacus, in vehicula arcu commodo ut. Aenean quis euismod lectus, vitae
          ornare ex. Vivamus posuere lectus at quam maximus semper. Fusce
          sagittis in est a consequat.
        </p>
      </section>
      <footer>Footer</footer>
    </article>
		    `,
	).WithDescription(
		`<p>Combine a <code>header</code> and a <code>footer</code> to give a card both a title and a place for related actions or metadata.</p>`,
	),
	NewExample(
		"Raised",
		`
    <article class="card raised">
      <header>Header</header>
      <section>
        <p>
          Vestibulum venenatis neque nec iaculis viverra. Proin a odio ex.
          Etiam eget rutrum nulla, id finibus tellus. Proin dignissim tortor
          lacus, in vehicula arcu commodo ut. Aenean quis euismod lectus, vitae
          ornare ex. Vivamus posuere lectus at quam maximus semper. Fusce
          sagittis in est a consequat.
        </p>
      </section>
      <footer>Footer</footer>
    </article>
		    `,
	).WithDescription(
		`<p>Add the <code>raised</code> class to visually raise the card with a shadow.</p>`,
	),
).WithDescription(
	`<p>An element that groups related information and actions about a single subject into a visually distinct, flexible container.</p>`,
)
