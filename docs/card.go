package docs

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

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
		P(Text("A minimal card that simply groups a single piece of content, such as a short summary or teaser, without any additional header or footer.")),
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
		P(Text("Add a "), Code(Text("header")), Text(" to a card to give the grouped content a title, so users can quickly identify what the card is about.")),
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
		P(Text("Add a "), Code(Text("footer")), Text(" to a card to hold supplementary information or actions, such as a call-to-action link, that relate to the card's content.")),
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
		P(Text("Combine a "), Code(Text("header")), Text(" and a "), Code(Text("footer")), Text(" to give a card both a title and a place for related actions or metadata.")),
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
		P(Text("Add the "), InlineCodeList("raised"), Text(" class to visually raise the card with a shadow.")),
	),
).WithDescription(
	P(Text("An element that groups related information and actions about a single subject into a visually distinct, flexible container.")),
)
