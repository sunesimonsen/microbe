package docs

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
	`<p>Lists let you present a set of related items, whether their order matters, as in an ordered list, doesn't matter, as in an unordered list, or each item pairs a term with its definition, as in a description list.</p>`,
)
