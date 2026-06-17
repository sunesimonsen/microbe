package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ListView() Node {
	return docpage(
		HGroup(H1(Text("List"))),
		example(
			"Unordered list",
			Ul(
				Li(Text("Gumbo beet greens"),
					Ul(
						Li(
							Text("Parsley shallot courgette"),
							Ul(
								Li(Text("Celery potato scallion")),
								Li(Text("Turnip cauliflower yarrow")),
								Li(Text("Corn amaranth salsify")),
							),
						),
						Li(Text("Pea horseradish azuki")),
						Li(Text("Chickweed okra coriander")),
					),
				),
				Li(Text("Grape kakadu plum")),
				Li(Text("Water spinach arugula")),
			),
		),
		example(
			"Ordered list",
			Ol(
				Li(Text("Gumbo beet greens"),
					Ol(
						Type("a"),
						Li(
							Text("Parsley shallot courgette"),
							Ol(
								Type("i"),
								Li(Text("Celery potato scallion")),
								Li(Text("Turnip cauliflower yarrow")),
								Li(Text("Corn amaranth salsify")),
							),
						),
						Li(Text("Pea horseradish azuki")),
						Li(Text("Chickweed okra coriander")),
					),
				),
				Li(Text("Grape kakadu plum")),
				Li(Text("Water spinach arugula")),
			),
		),
		example(
			"Mixed list",
			Ol(
				Li(Text("Gumbo beet greens"),
					Ol(
						Li(
							Text("Parsley shallot courgette"),
							Ul(
								Li(Text("Celery potato scallion")),
								Li(Text("Turnip cauliflower yarrow")),
								Li(Text("Corn amaranth salsify")),
							),
						),
						Li(Text("Pea horseradish azuki")),
						Li(Text("Chickweed okra coriander")),
					),
				),
				Li(Text("Grape kakadu plum")),
				Li(Text("Water spinach arugula")),
			),
		),
		example(
			"Description list",
			P(Text("Cryptids of Cornwall:")),
			Dl(
				Dt(Text("Beast of Bodmin")),
				Dd(Text("A large feline inhabiting Bodmin Moor.")),
				Dt(Text("Morgawr")),
				Dd(Text("A sea serpent.")),
				Dt(Text("Owlman")),
				Dd(Text("A giant owl-like creature.")),
			),
		),
	)
}
