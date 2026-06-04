package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ButtonView() Node {
	return Group([]Node{
		HGroup(
			H1(Text("Button")),
		),
		Div(
			Role("document"),
			example(
				"Styles",
				Class("grid"),
				Button(Class("solid"), Text("Solid Button")),
				Button(Class("outline"), Text("Outline Button")),
				Button(Class("ghost"), Text("Ghost Button")),
			),
			example(
				"Disabled",
				Class("grid"),
				Button(Class("solid"), Disabled(), Text("Solid Button")),
				Button(Class("outline"), Disabled(), Text("Outline Button")),
				Button(Class("ghost"), Disabled(), Text("Ghost Button")),
			),
			example(
				"Media",
				Class("grid"),
				Button(Class("outline"), gearIcon(), Text("Settings")),
				Button(Class("solid"), Text("Settings"), gearIcon()),
			),
			example(
				"Icon",
				Class("grid small"),
				Button(Class("solid icon"), burgerIcon()),
				Button(Class("outline icon"), burgerIcon()),
				Button(Class("ghost icon"), burgerIcon()),
			),
			example(
				"Types",
				Class("rows"),
				Div(
					Class("grid stretch"),
					Button(Class("solid"), Text("Button")),
					Button(Class("outline"), Text("Button")),
					Button(Class("ghost"), Text("Button")),
				),
				Div(
					Class("grid stretch"),
					Input(Class("solid"), Type("button"), Value("Input input")),
					Input(Class("outline"), Type("button"), Value("Input input")),
					Input(Class("ghost"), Type("button"), Value("Input input")),
				),
				Div(
					Class("grid stretch"),
					Input(Class("solid"), Type("submit"), Value("Summit input")),
					Input(Class("outline"), Type("submit"), Value("Summit input")),
					Input(Class("ghost"), Type("submit"), Value("Summit input")),
				),
				Div(
					Class("grid stretch"),
					Input(Class("solid"), Type("reset"), Value("Reset input")),
					Input(Class("outline"), Type("reset"), Value("Reset input")),
					Input(Class("ghost"), Type("reset"), Value("Reset input")),
				),
			),
		),
	})
}
