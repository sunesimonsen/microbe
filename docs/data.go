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
	).WithClass("grid"),
	NewExample(
		"Color",
		`
    <span class="tag" style="--neutral-hue: 20; --neutral-saturation: 80%;">NEW</span>
    <span class="tag" style="--neutral-hue: 40; --neutral-saturation: 80%;">NEW</span>
    <span class="tag" style="--neutral-hue: 100; --neutral-saturation: 80%;">NEW</span>
    `,
	).WithClass("grid"),
	NewExample(
		"Interactive",
		`
    <a class="tag" href="#interactive">Anchor</a>
    <button class="tag">Button</button>
    `,
	).WithClass("grid"),
)
