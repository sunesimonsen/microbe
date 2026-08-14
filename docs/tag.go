package docs

var TagPage = NewPage(
	"Tag",
	NewExample(
		"Default",
		`
    <span class="tag"><strong>Color</strong> black</span>
    <span class="tag"><strong>Size</strong> Medium</span>
    <span class="tag"><strong>Material</strong> Cotton</span>
		    `,
	).WithDescription(
		`<p>Tags are styled with the neutral color by default and can be <a href="#color">colorized using CSS variables</a>.</p>`,
	).WithClass("grid"),
	NewExample(
		"Color",
		`
    <span class="tag" style="--neutral-hue: 20; --neutral-saturation: 80%;">NEW</span>
    <span class="tag" style="--neutral-hue: 40; --neutral-saturation: 80%;">NEW</span>
    <span class="tag" style="--neutral-hue: 100; --neutral-saturation: 80%;">NEW</span>
		    `,
	).WithDescription(
		`<p>The color of the tags can be tweaked by setting the <code>--neutral-hue</code> and <code>--neutral-saturation</code> CSS variables.</p>`,
	).WithClass("grid small"),
	NewExample(
		"Interactive",
		`
    <a class="tag" href="#interactive">Anchor</a>
    <button class="tag">Button</button>
		    `,
	).WithDescription(
		`<p>Anchors and buttons are allowed to be styled as tags, this is useful for initiating a search for a given tag.</p>`,
	).WithClass("grid small"),
).WithDescription(
	`<p>Tags let users categorize content using a keyword.</p>`,
)
