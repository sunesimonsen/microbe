package views

import (
	"bytes"
	"log"
	"slices"
	"strings"

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
			A(Class("home"), Href("/"), Text("Microbe")),
		),
	)
}

func Nodes(nodes ...Node) Node {
	return Group(nodes)
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

func docsMenu(currentPath string) Node {
	return Aside(
		Class("menu"),

		Nav(
			Class("navlist"),
			navListSection("Content", currentPath, []pageRef{
				{href: "/typography", label: "Typography"},
				{href: "/list", label: "List"},
				{href: "/table", label: "Table"},
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
				{href: "/textarea", label: "Textarea"},
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

	return Nodes(
		Link(Rel("stylesheet"), Href("https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.11.1/styles/codepen-embed.min.css")),
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
	)
}

func example(name string, part ...Node) pageSection {
	buf := new(bytes.Buffer)
	err := Group(part).Render(buf)

	if err != nil {
		log.Fatal(err)
	}

	source := gohtml.Format(buf.String())
	source = strings.ReplaceAll(source, "&#39;", "'")

	return pageSection{
		name: name,
		content: Nodes(
			Article(
				Class("example"),
				ID(strcase.ToKebab(name)),
				H2(Text(name)),
				Article(
					Class("card"),
					Section(part...),
				),
				Section(
					Class("accordion"),
					Details(
						Class("hljs source "),
						Summary(Text("HTML")),
						Pre(Code(Data("highlight", "yes"), Class("language-html"), Text(source))),
					),
				),
			),
		),
	}
}

func IndexLayout(part Node) Node {
	return Page("Microbe",
		Main(
			Class("page-layout"),
			header(),
			part,
		),
	)
}

func DocsLayout(currentPath string, part Node) Node {
	return Page("Microbe",
		Main(
			Class("page-layout"),
			header(),
			docsMenu(currentPath),
			part,
		),
	)
}
