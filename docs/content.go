package docs

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

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
		P(Text("Use heading elements to establish a document's outline, from "), Code(Text("h1")), Text(" for the primary title down to "), Code(Text("h6")), Text(" for the least important sections, helping readers and assistive technology navigate the content.")),
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
		P(Text("Group a heading together with an introductory paragraph or tagline using "), Code(Text("hgroup")), Text(", so assistive technology announces them together as a single unit rather than as two disconnected pieces of content.")),
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
		P(Text("Microbe supports the standard inline text-level elements for indicating meaning within a sentence or phrase, such as emphasis, deletions, insertions, abbreviations, and keyboard input.")),
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
		P(Text("Use "), Code(Text("hr")), Text(" to signal a thematic break between paragraphs or sections, indicating that the topic or scene has shifted.")),
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
		P(Text("Use "), Code(Text("blockquote")), Text(" to quote text from another source, optionally combined with a "), Code(Text("footer")), Text(" and "), Code(Text("cite")), Text(" to attribute the quote to its author.")),
	),
	NewExample(
		"Code block",
		`
    <pre><code>console.log('Hello world!');</code></pre>
    `,
	).WithDescription(
		P(Text("Combine "), Code(Text("pre")), Text(" with "), Code(Text("code")), Text(" to display a multi-line, preformatted snippet of code or other text whose whitespace and line breaks must be preserved exactly as written.")),
	),
).WithDescription(
	P(Text("Microbe provides default styling for most typographic elements to allow easy creation of beautiful documents.")),
	P(
		Text("All styles are based on the "),
		Code(Text("em")),
		Text(" unit and will scale according to the surrounding font-size."),
	),
)

var ListPage = NewPage(
	"List",
	NewExample(
		"Unordered list",
		`
    <ul>
      <li>
        Gumbo beet greens
        <ul>
          <li>
            Parsley shallot courgette
            <ul>
              <li>Celery potato scallion</li>
              <li>Turnip cauliflower yarrow</li>
              <li>Corn amaranth salsify</li>
            </ul>
          </li>
          <li>Pea horseradish azuki</li>
          <li>Chickweed okra coriander</li>
        </ul>
      </li>
      <li>Grape kakadu plum</li>
      <li>Water spinach arugula</li>
    </ul>
    `,
	),
	NewExample(
		"Ordered list",
		`
    <ol>
      <li>
        Gumbo beet greens
        <ol type="a">
          <li>
            Parsley shallot courgette
            <ol type="i">
              <li>Celery potato scallion</li>
              <li>Turnip cauliflower yarrow</li>
              <li>Corn amaranth salsify</li>
            </ol>
          </li>
          <li>Pea horseradish azuki</li>
          <li>Chickweed okra coriander</li>
        </ol>
      </li>
      <li>Grape kakadu plum</li>
      <li>Water spinach arugula</li>
    </ol>
    `,
	),
	NewExample(
		"Mixed list",
		`
    <ol>
      <li>
        Gumbo beet greens
        <ol>
          <li>
            Parsley shallot courgette
            <ul>
              <li>Celery potato scallion</li>
              <li>Turnip cauliflower yarrow</li>
              <li>Corn amaranth salsify</li>
            </ul>
          </li>
          <li>Pea horseradish azuki</li>
          <li>Chickweed okra coriander</li>
        </ol>
      </li>
      <li>Grape kakadu plum</li>
      <li>Water spinach arugula</li>
    </ol>
    `,
	),
	NewExample(
		"Description list",
		`
    <p>Cryptids of Cornwall:</p>
    <dl>
      <dt>Beast of Bodmin</dt>
      <dd>A large feline inhabiting Bodmin Moor.</dd>
      <dt>Morgawr</dt>
      <dd>A sea serpent.</dd>
      <dt>Owlman</dt>
      <dd>A giant owl-like creature.</dd>
    </dl>
    `,
	),
).WithDescription(
	P(Text("Lists let you present a set of related items, whether their order matters, as in an ordered list, doesn't matter, as in an unordered list, or each item pairs a term with its definition, as in a description list.")),
)

var TablePage = NewPage(
	"Table",
	NewExample(
		"Default",
		`
    <table>
      <thead>
        <tr>
          <th scope="col">Planet</th>
          <th scope="col">Diameter (km)</th>
          <th scope="col">Distance to Sun (AU)</th>
          <th scope="col">Orbit (days)</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <th scope="row">Mercury</th>
          <td>4,880</td>
          <td>0.39</td>
          <td>88</td>
        </tr>
        <tr>
          <th scope="row">Venus</th>
          <td>12,104</td>
          <td>0.72</td>
          <td>225</td>
        </tr>
        <tr>
          <th scope="row">Earth</th>
          <td>12,742</td>
          <td>1.00</td>
          <td>365</td>
        </tr>
        <tr>
          <th scope="row">Mars</th>
          <td>6,779</td>
          <td>1.52</td>
          <td>687</td>
        </tr>
      </tbody>
      <tfoot>
        <th scope="row">Average</th>
        <td>9,126</td>
        <td>0.91</td>
        <td>341</td>
      </tfoot>
    </table>
    `,
	),
	NewExample(
		"Striped",
		`
    <table class="striped">
      <thead>
        <tr>
          <th scope="col">Planet</th>
          <th scope="col">Diameter (km)</th>
          <th scope="col">Distance to Sun (AU)</th>
          <th scope="col">Orbit (days)</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <th scope="row">Mercury</th>
          <td>4,880</td>
          <td>0.39</td>
          <td>88</td>
        </tr>
        <tr>
          <th scope="row">Venus</th>
          <td>12,104</td>
          <td>0.72</td>
          <td>225</td>
        </tr>
        <tr>
          <th scope="row">Earth</th>
          <td>12,742</td>
          <td>1.00</td>
          <td>365</td>
        </tr>
        <tr>
          <th scope="row">Mars</th>
          <td>6,779</td>
          <td>1.52</td>
          <td>687</td>
        </tr>
      </tbody>
      <tfoot>
        <th scope="row">Average</th>
        <td>9,126</td>
        <td>0.91</td>
        <td>341</td>
      </tfoot>
    </table>
    `,
	).WithDescription(
		P(Text("Add the "), InlineCodeList("striped"), Text(" class to color every second row to make the table more readable.")),
	),
).WithDescription(
	P(Text("Tables let you present tabular data organized into rows and columns, making it easy for users to compare related values at a glance.")),
)
