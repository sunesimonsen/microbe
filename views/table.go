package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func TableView() Node {
	return docpage(
		HGroup(
			H1(Text("Table")),
		),
		NewExample(
			"Default",
			Table(
				THead(
					Tr(
						Th(Scope("col"), Text("Planet")),
						Th(Scope("col"), Text("Diameter (km)")),
						Th(Scope("col"), Text("Distance to Sun (AU)")),
						Th(Scope("col"), Text("Orbit (days)")),
					),
				),
				TBody(
					Tr(
						Th(Scope("row"), Text("Mercury")),
						Td(Text("4,880")),
						Td(Text("0.39")),
						Td(Text("88")),
					),
					Tr(
						Th(Scope("row"), Text("Venus")),
						Td(Text("12,104")),
						Td(Text("0.72")),
						Td(Text("225")),
					),
					Tr(
						Th(Scope("row"), Text("Earth")),
						Td(Text("12,742")),
						Td(Text("1.00")),
						Td(Text("365")),
					),
					Tr(
						Th(Scope("row"), Text("Mars")),
						Td(Text("6,779")),
						Td(Text("1.52")),
						Td(Text("687")),
					),
				),
				TFoot(
					Th(Scope("row"), Text("Average")),
					Td(Text("9,126")),
					Td(Text("0.91")),
					Td(Text("341")),
				),
			),
		),
		NewExample(
			"Striped",
			Table(
				Class("striped"),
				THead(
					Tr(
						Th(Scope("col"), Text("Planet")),
						Th(Scope("col"), Text("Diameter (km)")),
						Th(Scope("col"), Text("Distance to Sun (AU)")),
						Th(Scope("col"), Text("Orbit (days)")),
					),
				),
				TBody(
					Tr(
						Th(Scope("row"), Text("Mercury")),
						Td(Text("4,880")),
						Td(Text("0.39")),
						Td(Text("88")),
					),
					Tr(
						Th(Scope("row"), Text("Venus")),
						Td(Text("12,104")),
						Td(Text("0.72")),
						Td(Text("225")),
					),
					Tr(
						Th(Scope("row"), Text("Earth")),
						Td(Text("12,742")),
						Td(Text("1.00")),
						Td(Text("365")),
					),
					Tr(
						Th(Scope("row"), Text("Mars")),
						Td(Text("6,779")),
						Td(Text("1.52")),
						Td(Text("687")),
					),
				),
				TFoot(
					Th(Scope("row"), Text("Average")),
					Td(Text("9,126")),
					Td(Text("0.91")),
					Td(Text("341")),
				),
			),
		),
	)
}
