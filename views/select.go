package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func SelectView() Node {
	return docpage(
		HGroup(H1(Text("Select"))),
		NewExample(
			"Example",
			Nodes(
				Label(
					Text("Cuisine"),
					Select(
						Aria("label", "Select you favorite cuisine..."),
						Option(Selected(), Disabled(), Value(""), Text("Select your favorite cuisine...")),
						Option(Value("0"), Text("Italian")),
						Option(Value("1"), Text("Japanese")),
						Option(Value("2"), Text("Indian")),
						Option(Value("3"), Text("Thai")),
						Option(Value("4"), Text("French")),
					),
				),
				Label(
					Text("Snacks"),
					Select(
						Aria("label", "Select you favorite snacks..."),
						Multiple(),
						Option(Disabled(), Value(""), Text("Select your favorite snacks...")),
						Option(Value("cheese"), Text("Cheese")),
						Option(Value("fruits"), Text("Fruits")),
						Option(Value("nuts"), Text("Nuts")),
						Option(Value("chocolate"), Text("Chocolate")),
						Option(Value("crackers"), Text("Crackers")),
					),
				),
			),
		),
		NewExample(
			"Single value",
			Label(
				Text("Select a number"),
				Select(
					Aria("label", "Select a number"),
					Option(Value("one"), Text("One")),
					Option(Value("two"), Text("Two")),
					Option(Value("three"), Text("Three")),
				),
			),
		),
		NewExample(
			"Multiple values",
			Label(
				Text("Select colors"),
				Select(
					Aria("label", "Select colors"),
					Multiple(),
					Option(Value("blue"), Text("Blue")),
					Option(Value("green"), Text("Green")),
					Option(Value("orange"), Text("Orange")),
					Option(Value("purple"), Text("Purple")),
					Option(Value("red"), Text("Red")),
					Option(Value("yellow"), Text("Yellow")),
				),
			),
		),
		NewExample(
			"Disabled",
			Nodes(
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
		),
		NewExample(
			"Hint",
			Nodes(
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
		),
		NewExample(
			"Validation",
			Nodes(
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
	)
}
