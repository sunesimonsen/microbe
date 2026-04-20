package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Buttons() Node {
	return Group([]Node{
		HGroup(
			H1(Text("Buttons")),
		),
		Div(
			Attr("role", "document"),
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
		),
	})
}
