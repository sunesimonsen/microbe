package docs

import (
	"fmt"
	"strconv"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
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

var lightnessSteps = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

func colorSample(c color, lightness int) Node {
	return Button(
		Class("color-sample"),
		Data("hue", strconv.Itoa(c.hue)),
		Data("saturation", strconv.Itoa(c.saturation)),
		Data("lightness", strconv.Itoa(lightness)),
		Style(fmt.Sprintf("--lightness: var(--lightness-%d); --hue: %d; --saturation: %d%%", lightness, c.hue, c.saturation)),
	)
}

func colorTable(name string, palettes []color) PageSection {
	rows := []Node{}

	for _, palette := range palettes {
		rowCells := []Node{}
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
			Style(fmt.Sprintf("--columns: %d", len(lightnessSteps))),
			Group(rows),
		),
	)
}

var PalettePage = NewPage(
	"Palette",
	colorTable("Standard", createStandardPalettes()),
	colorTable("Muted", createCustomPalettes()),
).WithDescription(
	`
  <p>Microbe defines 12 lightness levels that can be combine with any hue and saturation to create a color. The lightness levels allows the colors to work both in light and dark modes.</p>
  <p>Click on any of the colors to copy the definition to the clipboard.</p>
  `,
)
