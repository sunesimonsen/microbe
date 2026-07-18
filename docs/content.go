package docs

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var TypographyPage = NewPage(
	"Typography",
	HGroup(
		H1(Text("Typography")),
		P(Text("Microbe provides default styling for most typographic elements to allow easy creation of beautiful documents.")),
		P(
			Text("All styles are based on the HTML document's font-size ("),
			Code(Text("rem")),
			Text(") and will scale according to the users stylesheet and the screen size."),
		),
	),
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
	),
	NewExample(
		"Heading group",
		`
    <hgroup>
      <h2>Get inspired with CSS</h2>
      <p>How to use CSS to add glam to your Website?</p>
    </hgroup>
    `,
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
	),
	NewExample(
		"Code block",
		`
    <pre><code>console.log('Hello world!');</code></pre>
    `,
	),
)

var ListPage = NewPage(
	"List",
	HGroup(H1(Text("List"))),
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
)

var TablePage = NewPage(
	"Table",
	HGroup(H1(Text("Table"))),
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
	),
)
