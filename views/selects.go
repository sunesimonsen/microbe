package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Selects() Node {
	return Group([]Node{
		HGroup(
			H1(Text("Selects")),
		),
		Div(
			Attr("role", "document"),
			example(
				"Example",
				Label(
					Text("Cuisine"),
					Select(
						Name("favorite-cuisine"),
						Aria("label", "Select you favorite cuisine..."),
						Option(Selected(), Disabled(), Value(""), Text("Select your favorite cuisine...")),
						Option(Text("Italian")),
						Option(Text("Japanese")),
						Option(Text("Indian")),
						Option(Text("Thai")),
						Option(Text("French")),
					),
				),
				Label(
					Text("Snacks"),
					Select(
						Name("favorite-snacks"),
						Aria("label", "Select you favorite snacks..."),
						Multiple(),
						Option(Disabled(), Value(""), Text("Select your favorite snacks...")),
						Option(Text("Cheese")),
						Option(Text("Fruits")),
						Option(Text("Nuts")),
						Option(Text("Chocolate")),
						Option(Text("Crackers")),
					),
				),
			),
			example(
				"Disabled",
				Label(
					Text("Cuisine"),
					Select(
						Aria("label", "Select you favorite cuisine..."),
						Disabled(),
						Option(Selected(), Disabled(), Value(""), Text("Select your favorite cuisine...")),
						Option(Text("Italian")),
						Option(Text("Japanese")),
						Option(Text("Indian")),
						Option(Text("Thai")),
						Option(Text("French")),
					),
				),
				Label(
					Text("Snacks"),
					Select(
						Name("favorite-snacks"),
						Aria("label", "Select you favorite snacks..."),
						Multiple(),
						Disabled(),
						Option(Disabled(), Value(""), Text("Select your favorite snacks...")),
						Option(Text("Cheese")),
						Option(Text("Fruits")),
						Option(Text("Nuts")),
						Option(Text("Chocolate")),
						Option(Text("Crackers")),
					),
				),
			),
			example(
				"Hint",
				Label(
					Text("Cuisine"),
					Select(
						Aria("label", "Select you favorite cuisine..."),
						Aria("describedby", "favorite-cuisine-hint"),
						Option(Selected(), Disabled(), Value(""), Text("Select your favorite cuisine...")),
						Option(Text("Italian")),
						Option(Text("Japanese")),
						Option(Text("Indian")),
						Option(Text("Thai")),
						Option(Text("French")),
					),
				),
				Small(ID("favorite-cuisine-hint"), Text("Select you favorite cuisine")),
			),
			example(
				"Validation",
				Class("rows"),
				Div(
					Select(
						Name("pizza-topping"),
						Aria("label", "Select your favorite pizza topping..."),
						Aria("invalid", "false"),
						Aria("describedby", "success-hint"),
						Option(Disabled(), Text("Select your favorite pizza topping...")),
						Option(Selected(), Text("Pepperoni")),
						Option(Text("Mushrooms")),
						Option(Text("Onions")),
						Option(Text("Green Peppers")),
						Option(Text("Olives")),
					),
					Small(ID("success-hint"), Text("Great choice!")),
				),
				Div(
					Select(
						Name("pizza-topping"),
						Aria("label", "Select your favorite pizza topping..."),
						Aria("invalid", "true"),
						Aria("describedby", "failure-hint"),
						Option(Disabled(), Text("Select your favorite pizza topping...")),
						Option(Text("Pepperoni")),
						Option(Text("Mushrooms")),
						Option(Text("Onions")),
						Option(Text("Green Peppers")),
						Option(Text("Olives")),
					),
					Small(ID("failure-hint"), Text("Please provide a valid value!")),
				),
			),
		),
	})
}
