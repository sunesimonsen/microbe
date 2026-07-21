package docs

import (
	"fmt"

	"github.com/iancoleman/strcase"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var AccentColorPage = NewPage(
	"Accent color",
	NewPageSection(
		"Settings",
		Section(
			ID("settings"),
			Label(
				Text("Hue"),
				Input(
					Attr("type", "range"),
					Attr("value", "240"),
					Attr("min", "0"),
					Attr("max", "359"),
					Attr("id", "color-range"),
				)),
			Label(
				Text("Saturation"),
				Input(
					Attr("type", "range"),
					Attr("value", "70"),
					Attr("min", "0"),
					Attr("max", "100"),
					Attr("id", "saturation-range"),
				)),
		),
	),
	NewExample(
		"Example",
		`
    <form id="theming-example-form">
      <fieldset>
        <label>
          Name
          <input name="name" placeholder="Name" autocomplete="name">
        </label>
        <label>
          Email
          <input name="email" placeholder="Email" autocomplete="email" aria-describedby="email-hint">
          <small id="email-hint">
            We’ll never share your email with anyone else.
          </small>
        </label>
        <label>
          <input type="checkbox" name="newsletter" aria-describedby="newsletter-hint" checked>
          Newsletter
        </label>
        <small id="newsletter-hint">
          We will send you a news letter every week
        </small>
        <label>
          Spelling proficiency (0 - 10)
          <input type="range" value="5" min="0" max="10">
        </label>
      </fieldset>
      <div class="actions">
        <button class="outline" type="reset">
          Reset
        </button>
        <button class="solid" type="submit">
          Submit
        </button>
      </div>
    </form>
    `,
	),
)

type color struct {
	hue        int
	saturation int
}

func createStandardPalettes() []color {
	palettes := []color{}

	for hue := 0; hue < 360; hue += 10 {
		if hue < 60 || 180 < hue {
			palettes = append(palettes, color{hue: hue, saturation: 90})
		} else {
			palettes = append(palettes, color{hue: hue, saturation: 70})
		}
	}

	return palettes
}

func createCustomPalettes() []color {
	return []color{
		{hue: 45, saturation: 40},
		{hue: 45, saturation: 30},
		{hue: 45, saturation: 20},
		{hue: 45, saturation: 10},
		{hue: 0, saturation: 0},
		{hue: 220, saturation: 10},
		{hue: 220, saturation: 20},
		{hue: 220, saturation: 30},
		{hue: 220, saturation: 40},
	}
}

var lightnessSteps = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

func colorSample(c color, lightness int) Node {
	return Button(
		Class("color-sample"),
		Style(fmt.Sprintf("--lightness: var(--lightness-%d); --hue: %d; --saturation: %d%%", lightness, c.hue, c.saturation)),
	)
}

func colorTable(name string, palettes []color) PageSection {
	headerCells := []Node{Div(Class("color-grid-header"))}
	for _, lightness := range lightnessSteps {
		headerCells = append(headerCells, Div(Class("color-grid-header"), Text(fmt.Sprintf("%d", lightness))))
	}

	rows := []Node{Group(headerCells)}

	for _, palette := range palettes {
		rowCells := []Node{Div(Class("color-grid-header"), Text(fmt.Sprintf("%d", palette.hue)))}
		for _, lightness := range lightnessSteps {
			rowCells = append(rowCells, colorSample(palette, lightness))
		}
		rows = append(rows, Group(rowCells))
	}

	return NewPageSection(
		name,
		H2(ID(strcase.ToKebab(name)), Text(name)),
		Section(
			Class("color-grid"),
			Style(fmt.Sprintf("--columns: %d", len(lightnessSteps)+1)),
			Group(rows),
		),
	)
}

var ColorsPage = NewPage(
	"Colors",
	colorTable("Standard colors", createStandardPalettes()),
	colorTable("Muted colors", createCustomPalettes()),
)
