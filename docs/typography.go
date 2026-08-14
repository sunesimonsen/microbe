package docs

var TypographyPage = NewPage(
	"Typography",
	NewExample(
		"Headings",
		`
    <h1>Heading 1</h1>
    <p>
      Vestibulum venenatis neque nec iaculis viverra. Proin a odio ex.
    </p>
    <h2>Heading 2</h2>
    <p>
      Phasellus nec luctus dolor. Curabitur id facilisis diam.
    </p>
    <h3>Heading 3</h3>
    <p>
      Vivamus feugiat tempor tellus, vel consequat turpis gravida non.
    </p>
    <h4>Heading 4</h4>
    <p>
      Quisque tristique lobortis ligula id tempus.
    </p>
    <h5>Heading 5</h5>
    <p>
      Sed aliquet velit mauris, vel interdum diam mattis et.
    </p>
    <h6>Heading 6</h6>
    <p>
      Donec in lorem imperdiet, eleifend turpis eget, congue velit.
    </p>
		    `,
	).WithDescription(
		`<p>Use heading elements to establish a document's outline, from <code>h1</code> for the primary title down to <code>h6</code> for the least important sections, helping readers and assistive technology navigate the content.</p>`,
	),
	NewExample(
		"Heading group",
		`
    <hgroup>
      <h2>Get inspired with CSS</h2>
      <p>How to use CSS to add glam to your Website?</p>
    </hgroup>
		    `,
	).WithDescription(
		`<p>Group a heading together with an introductory paragraph or tagline using <code>hgroup</code>, so assistive technology announces them together as a single unit rather than as two disconnected pieces of content.</p>`,
	),
	NewExample(
		"Inline text",
		`
    <abbr>Abbreviation</abbr>
    <cite>Citation</cite>
    <code>Code</code>
    <del>Deleted</del>
    <dfn>Definition</dfn>
    <em>Emphasised</em>
    <i>Idiomatic</i>
    <ins>Inserted</ins>
    <kbd>Ctrl + S</kbd>
    <mark>Highlighted</mark>
    <s>Strikethrough</s>
    <samp>Samp</samp>
    <small>Small</small>
    <span>X<sub>sub</sub></span>
    <span>X<sup>sup</sup></span>
    <strong>Strong</strong>
    <u>Underlined</u>
    <var>Var</var>
		    `,
	).WithDescription(
		`<p>Microbe supports the standard inline text-level elements for indicating meaning within a sentence or phrase, such as emphasis, deletions, insertions, abbreviations, and keyboard input.</p>`,
	).WithClass("grid"),
	NewExample(
		"Horizontal ruler",
		`
    <p>
      Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla ullamcorper
      blandit ultricies. Etiam non suscipit felis. Orci varius natoque penatibus et
      magnis dis parturient montes, nascetur ridiculus mus. Proin feugiat purus
      hendrerit sapien condimentum, eu facilisis ex eleifend. Sed lobortis est a
      urna accumsan vestibulum. Curabitur iaculis sem lacus, id molestie eros
      ultrices eu. Nullam justo nulla, sollicitudin sed mi pretium, consequat
      varius erat. Phasellus sed eros dictum, congue enim in, dictum ante. Sed
      condimentum ut elit eu vehicula.
    </p>
    <hr>
    <p>
      Fusce congue nec massa id eleifend. Donec id luctus ligula. In pellentesque
      diam eu magna interdum, a tristique arcu dignissim. In eu lacinia nisl.
      Phasellus eu nisi vitae enim aliquam rutrum. Suspendisse fringilla tortor et
      tincidunt vulputate. Nam a urna a purus ornare gravida non sit amet risus.
      Morbi in justo quis velit elementum fermentum. Etiam tristique diam nunc,
      quis suscipit velit suscipit non.
    </p>
		    `,
	).WithDescription(
		`<p>Use <code>hr</code> to signal a thematic break between paragraphs or sections, indicating that the topic or scene has shifted.</p>`,
	),
	NewExample(
		"Blockquote",
		`
    <blockquote>
      "Be the change that you wish to see in the world."
      <footer>
        <cite>— Mahatma Gandhi</cite>
      </footer>
    </blockquote>
		    `,
	).WithDescription(
		`<p>Use <code>blockquote</code> to quote text from another source, optionally combined with a <code>footer</code> and <code>cite</code> to attribute the quote to its author.</p>`,
	),
	NewExample(
		"Code block",
		`
    <pre><code>console.log('Hello world!');</code></pre>
		    `,
	).WithDescription(
		`<p>Combine <code>pre</code> with <code>code</code> to display a multi-line, preformatted snippet of code or other text whose whitespace and line breaks must be preserved exactly as written.</p>`,
	),
).WithDescription(
	`<p>Microbe provides default styling for most typographic elements to allow easy creation of beautiful documents.</p><p>All styles are based on the <code>em</code> unit and will scale according to the surrounding font-size.</p>`,
)
