package views

import (
	"slices"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func header() Node {
	return Header(
		Nav(
			Button(
				Class("menu-toggle ghost icon"),
				burgerIcon(),
			),
			Text("Microbe"),
		),
	)
}

type pageRef struct {
	href  string
	label string
}

func navListSection(label string, currentPath string, pageRefs ...pageRef) Node {
	hasActivePage := slices.ContainsFunc(pageRefs, func(pr pageRef) bool {
		return pr.href == currentPath
	})

	return Details(
		If(hasActivePage, Open()),
		Summary(Text(label)),
		Ul(
			Map(pageRefs, func(pr pageRef) Node {
				return Li(A(
					If(pr.href == currentPath, Aria("current", "page")),
					Href(pr.href),
					Text(pr.label),
				))
			}),
		),
	)
}

func menu(currentPath string) Node {
	return Aside(
		Class("menu"),

		Nav(
			Class("navlist"),
			navListSection("Content", currentPath, []pageRef{
				{href: "/typography", label: "Typography"},
			}...),
			navListSection("Navigation", currentPath, []pageRef{
				{href: "/anchor", label: "Anchor"},
				{href: "/navlist", label: "Navlist"},
			}...),
			navListSection("Layout", currentPath, []pageRef{
				{href: "/spacing", label: "Spacing"},
			}...),
			navListSection("Forms", currentPath, []pageRef{
				{href: "/button", label: "Button"},
				{href: "/checkbox", label: "Checkbox"},
				{href: "/input", label: "Input"},
				{href: "/radio", label: "Radio"},
				{href: "/range", label: "Range"},
				{href: "/select", label: "Select"},
				{href: "/switch", label: "Switch"},
			}...),
			navListSection("Loaders", currentPath, []pageRef{
				{href: "/progress", label: "Progress"},
			}...),
			navListSection("Components", currentPath, []pageRef{
				{href: "/card", label: "Card"},
				{href: "/dialog", label: "Dialog"},
			}...),
			navListSection("Theming", currentPath, []pageRef{
				{href: "/accent-color", label: "Accent color"},
				{href: "/colors", label: "Colors"},
			}...),
		),
	)
}

func example(description string, part ...Node) Node {
	return Article(
		Class("example card"),
		Header(Text(description)),
		Section(
			part...,
		),
	)
}

func IndexLayout(part Node) Node {
	return Page("Microbe",
		Main(
			Class("standard-layout"),
			header(),
			part,
		),
	)
}

func DocsLayout(currentPath string, part Node) Node {
	return Page("Microbe",
		Main(
			Class("standard-layout"),
			header(),
			menu(currentPath),
			part,
		),
	)
}
