package views

import (
	"bytes"
	"log"
	"slices"

	"github.com/iancoleman/strcase"
	"github.com/yosssi/gohtml"
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
		Name("index"),
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
				{href: "/accordion", label: "Accordion"},
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

type pageSection struct {
	name    string
	content Node
}

func (s pageSection) fragment() string {
	return strcase.ToKebab(s.name)
}

func docpage(header Node, sections ...pageSection) Node {
	content := []Node{
		Role("document"),
		Map(sections, func(s pageSection) Node {
			return s.content
		}),
	}

	return Group([]Node{
		Link(Rel("stylesheet"), Href("https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.11.1/styles/github-dark-dimmed.min.css")),
		Script(Src("https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.11.1/highlight.min.js")),
		HGroup(header),
		Div(Group(content)),
		Aside(
			Class("toc"),
			Nav(
				Class("navlist"),
				Details(
					Open(),
					Summary(Text("Content")),
					Name("toc"),
					Ul(
						Map(sections, func(section pageSection) Node {
							return Li(A(Href("#"+section.fragment()), Text(section.name)))
						})),
				),
			),
		),
	})
}

func example(name string, part ...Node) pageSection {
	buf := new(bytes.Buffer)
	err := Group(part).Render(buf)

	if err != nil {
		log.Fatal(err)
	}

	source := gohtml.Format(buf.String())

	return pageSection{
		name: name,
		content: Group([]Node{
			Article(
				ID(strcase.ToKebab(name)),
				Class("example card"),
				Header(Text(name)),
				Section(part...),
				Pre(Code(Class("language-html"), Text(source))),
			),
		}),
	}
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
