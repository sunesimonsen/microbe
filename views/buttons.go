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
			Div(
				Class("buttons-row"),
				Button(Class("solid"), Text("Solid Button")),
				Button(Class("outline"), Text("Outline Button")),
				Button(Class("ghost"), Text("Ghost Button")),
			),
			Div(
				Class("buttons-row"),
				Button(Class("solid"), Text("Solid Button")),
				Button(Class("outline"), Text("Outline Button")),
				Button(Class("ghost"), Text("Ghost Button")),
			),
			Div(
				Class("buttons-row"),
				Button(Class("solid"), Disabled(), Text("Solid Button")),
				Button(Class("outline"), Disabled(), Text("Outline Button")),
				Button(Class("ghost"), Disabled(), Text("Ghost Button")),
			),
			Div(
				Class("buttons-row"),
				Button(Class("solid"), gearIcon(), Text(" Solid Button")),
				Button(Class("outline"), gearIcon(), Text(" Outline Button")),
				Button(Class("ghost"), gearIcon(), Text(" Ghost Button")),
			),
			Div(
				Class("buttons-row"),
				Button(Class("solid"), Text("Solid Button "), gearIcon()),
				Button(Class("outline"), Text("Outline Button "), gearIcon()),
				Button(Class("ghost"), Text("Ghost Button "), gearIcon()),
			),
			Div(
				Class("buttons-row"),
				Button(Class("solid icon"), burgerIcon()),
				Button(Class("outline icon"), burgerIcon()),
				Button(Class("ghost icon"), burgerIcon()),
			),
		),
	})
}
