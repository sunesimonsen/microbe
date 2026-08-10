package docs

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

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
		P(Text("Tags are styled with the neutral color by default and can be "), A(Href("#color"), Text("colorized using CSS variables")), Text(".")),
	).WithClass("grid"),
	NewExample(
		"Color",
		`
    <span class="tag" style="--neutral-hue: 20; --neutral-saturation: 80%;">NEW</span>
    <span class="tag" style="--neutral-hue: 40; --neutral-saturation: 80%;">NEW</span>
    <span class="tag" style="--neutral-hue: 100; --neutral-saturation: 80%;">NEW</span>
    `,
	).WithDescription(
		P(Text("The color of the tags can be tweaked by setting the "), InlineCodeList("--neutral-hue", "--neutral-saturation"), Text("CSS variables.")),
	).WithClass("grid small"),
	NewExample(
		"Interactive",
		`
    <a class="tag" href="#interactive">Anchor</a>
    <button class="tag">Button</button>
    `,
	).WithDescription(
		P(Text("Anchors and buttons are allowed to be styled as tags, this is useful for initiating a search for a given tag.")),
	).WithClass("grid small"),
).WithDescription(
	P(Text("Tags are UI elements that indicate a segmentation of data. Put to a tag on an entity to indicate that it belongs to a category.")),
)
