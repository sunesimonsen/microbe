package docs

import (
	"fmt"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var AccentColorPage = NewPage(
	"Accent color",
	NewStaticPageSection(
		"Playground",
		Article(
			H2(Text("Playground")),
			Div(
				Class("card raised"),
				Header(
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
				Section(
					Form(
						ID("theming-example-form"),
						FieldSet(
							Label(
								Text("Name"),
								Input(Placeholder("Name"), AutoComplete("name")),
							),
							Label(
								Text("Email"),
								Input(Placeholder("Email"), AutoComplete("email"), Aria("describedby", "email-hint")),
								Small(ID("email-hint"), Text("We'll never share your email with anyone else.")),
							),
							Label(
								Input(Type("checkbox"), Checked(), Aria("describedby", "newsletter-hint")),
								Text("Newsletter"),
							),
							Small(ID("newsletter-hint"), Text("We will send you a newsletter every week.")),
						),
						Div(
							Class("actions"),
							Button(Class("outline"), Type("reset"), Text("Reset")),
							Button(Class("solid"), Type("submit"), Text("Submit")),
						),
					),
				),
			),
		),
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

	return NewStaticPageSection(
		name,
		H2(Text(name)),
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
