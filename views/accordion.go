package views

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func AccordionView() Node {
	return docpage(
		HGroup(
			H1(Text("Accordion")),
			P(Text("An element that organizes content into a vertically stacked list of collapsible sections. Users can click or tap a section's header to expand it and reveal detailed information, or collapse it to hide the content and reduce scrolling.")),
		),
		NewExample(
			"Default",
			Nodes(
				P(Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla ullamcorper blandit ultricies. Etiam non suscipit felis. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Proin feugiat purus hendrerit sapien condimentum, eu facilisis ex eleifend. Sed lobortis est a urna accumsan vestibulum. Curabitur iaculis sem lacus, id molestie eros ultrices eu. Nullam justo nulla, sollicitudin sed mi pretium, consequat varius erat. Phasellus sed eros dictum, congue enim in, dictum ante. Sed condimentum ut elit eu vehicula.")),
				Div(
					Class("accordion"),
					Details(
						Summary(Text("Section 1")),
						P(Text("Fusce congue nec massa id eleifend. Donec id luctus ligula. In pellentesque diam eu magna interdum, a tristique arcu dignissim. In eu lacinia nisl. Phasellus eu nisi vitae enim aliquam rutrum. Suspendisse fringilla tortor et tincidunt vulputate. Nam a urna a purus ornare gravida non sit amet risus. Morbi in justo quis velit elementum fermentum. Etiam tristique diam nunc, quis suscipit velit suscipit non.")),
					),
					Details(
						Summary(Text("Section 2")),
						P(Text("Nam at mollis ante, non rutrum enim. Nulla facilisi. Nunc maximus diam a dui tincidunt vulputate. Phasellus in suscipit augue. In quis elementum enim. Mauris at porttitor libero. Sed vel tortor sit amet felis porttitor dapibus et id neque.")),
					),
					Details(
						Summary(Text("Section 3")),
						P(Text("Duis facilisis eros porttitor mauris sollicitudin dapibus. Mauris mollis turpis nec ligula pulvinar finibus. Proin molestie varius nisl. Ut et ex lobortis, auctor nunc tincidunt, aliquam sem. Nunc sodales aliquam magna. Proin rutrum sed erat accumsan elementum. Mauris posuere augue non orci maximus, sed sagittis elit mattis. Proin mattis, leo at luctus ultrices, nulla dui bibendum ipsum, sed venenatis ex ipsum a erat. Ut libero dui, viverra in purus eget, efficitur pellentesque justo.")),
					),
					Details(
						Summary(Text("Section 4")),
						P(Text("Suspendisse ultricies tristique elit eget suscipit. Sed eu odio eu diam convallis venenatis sit amet quis elit. Aenean mi diam, gravida eu nulla eget, egestas ultricies velit. Pellentesque et metus fringilla, lobortis felis id, suscipit nunc. Nunc interdum placerat dui, quis varius turpis tempor in. Maecenas faucibus tellus non lacus pulvinar, a malesuada urna luctus. Phasellus fermentum porta dolor vel hendrerit. Donec facilisis maximus ipsum congue mattis. Nullam justo arcu, aliquet id turpis sed, maximus laoreet dui. Nam nec elit accumsan, facilisis lorem ac, ullamcorper arcu. Interdum et malesuada fames ac ante ipsum primis in faucibus. Nullam quis rutrum libero, a sollicitudin magna.")),
					),
				),
				P(Text("Pellentesque at sodales nisl, in eleifend erat. Pellentesque porttitor lacinia nibh, a dictum orci condimentum sed. Quisque eu maximus leo. Vestibulum elit arcu, molestie eget posuere vel, dapibus pharetra elit. Aenean eget turpis et turpis ullamcorper pretium. Praesent rutrum sollicitudin leo quis tincidunt. Aliquam erat volutpat. Ut viverra, dui non placerat blandit, urna justo gravida erat, sit amet semper dolor elit sed metus. Mauris volutpat purus vitae sodales eleifend. Cras consequat scelerisque elit, sed mattis tortor lacinia non.")),
			),
		),
		NewExample(
			"Single panel",
			Nodes(
				P(Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla ullamcorper blandit ultricies. Etiam non suscipit felis. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Proin feugiat purus hendrerit sapien condimentum, eu facilisis ex eleifend. Sed lobortis est a urna accumsan vestibulum. Curabitur iaculis sem lacus, id molestie eros ultrices eu. Nullam justo nulla, sollicitudin sed mi pretium, consequat varius erat. Phasellus sed eros dictum, congue enim in, dictum ante. Sed condimentum ut elit eu vehicula.")),
				Div(
					Class("accordion"),
					Details(
						Name("single"),
						Summary(Text("Section 1")),
						P(Text("Fusce congue nec massa id eleifend. Donec id luctus ligula. In pellentesque diam eu magna interdum, a tristique arcu dignissim. In eu lacinia nisl. Phasellus eu nisi vitae enim aliquam rutrum. Suspendisse fringilla tortor et tincidunt vulputate. Nam a urna a purus ornare gravida non sit amet risus. Morbi in justo quis velit elementum fermentum. Etiam tristique diam nunc, quis suscipit velit suscipit non.")),
					),
					Details(
						Name("single"),
						Summary(Text("Section 2")),
						P(Text("Nam at mollis ante, non rutrum enim. Nulla facilisi. Nunc maximus diam a dui tincidunt vulputate. Phasellus in suscipit augue. In quis elementum enim. Mauris at porttitor libero. Sed vel tortor sit amet felis porttitor dapibus et id neque.")),
					),
					Details(
						Name("single"),
						Summary(Text("Section 3")),
						P(Text("Duis facilisis eros porttitor mauris sollicitudin dapibus. Mauris mollis turpis nec ligula pulvinar finibus. Proin molestie varius nisl. Ut et ex lobortis, auctor nunc tincidunt, aliquam sem. Nunc sodales aliquam magna. Proin rutrum sed erat accumsan elementum. Mauris posuere augue non orci maximus, sed sagittis elit mattis. Proin mattis, leo at luctus ultrices, nulla dui bibendum ipsum, sed venenatis ex ipsum a erat. Ut libero dui, viverra in purus eget, efficitur pellentesque justo.")),
					),
					Details(
						Name("single"),
						Summary(Text("Section 4")),
						P(Text("Suspendisse ultricies tristique elit eget suscipit. Sed eu odio eu diam convallis venenatis sit amet quis elit. Aenean mi diam, gravida eu nulla eget, egestas ultricies velit. Pellentesque et metus fringilla, lobortis felis id, suscipit nunc. Nunc interdum placerat dui, quis varius turpis tempor in. Maecenas faucibus tellus non lacus pulvinar, a malesuada urna luctus. Phasellus fermentum porta dolor vel hendrerit. Donec facilisis maximus ipsum congue mattis. Nullam justo arcu, aliquet id turpis sed, maximus laoreet dui. Nam nec elit accumsan, facilisis lorem ac, ullamcorper arcu. Interdum et malesuada fames ac ante ipsum primis in faucibus. Nullam quis rutrum libero, a sollicitudin magna.")),
					),
				),
				P(Text("Pellentesque at sodales nisl, in eleifend erat. Pellentesque porttitor lacinia nibh, a dictum orci condimentum sed. Quisque eu maximus leo. Vestibulum elit arcu, molestie eget posuere vel, dapibus pharetra elit. Aenean eget turpis et turpis ullamcorper pretium. Praesent rutrum sollicitudin leo quis tincidunt. Aliquam erat volutpat. Ut viverra, dui non placerat blandit, urna justo gravida erat, sit amet semper dolor elit sed metus. Mauris volutpat purus vitae sodales eleifend. Cras consequat scelerisque elit, sed mattis tortor lacinia non.")),
			),
		),
	)
}
