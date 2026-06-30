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
			navListSection("Getting started", currentPath, []pageRef{
				{href: "/", label: "Introduction"},
			}...),
			navListSection("Theming", currentPath, []pageRef{
				{href: "/accent-color", label: "Accent color"},
				{href: "/colors", label: "Colors"},
			}...),
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
			navListSection("Popups", currentPath, []pageRef{
				{href: "/dialog", label: "Dialog"},
			}...),
			navListSection("Components", currentPath, []pageRef{
				{href: "/accordion", label: "Accordion"},
				{href: "/card", label: "Card"},
			}...),
		),
	)
}

type PageSection struct {
	name    string
	content Node
}

func NewPageSection(name string, content ...Node) PageSection {
	return PageSection{name: name, content: Group(content)}
}

func (s PageSection) Name() string {
	return s.name
}

func (s PageSection) Anchor() Node {
	return A(Href("#"+strcase.ToKebab(s.name)), Text(s.name))
}

func (s PageSection) Content() Node {
	return s.content
}

type IndexedContent interface {
	Name() string
	Anchor() Node
	Content() Node
}

func docpage(header Node, sections ...IndexedContent) Node {
	content := []Node{
		Role("document"),
		Map(sections, func(s IndexedContent) Node {
			return s.Content()
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
						Map(sections, func(s IndexedContent) Node {
							return Li(s.Anchor())
						})),
				),
			),
		),
	)
}

type ExampleOption func(example *Example)

func WithDescription(description Node) ExampleOption {
	return func(example *Example) {
		example.description = description
	}
}

type Example struct {
	PageSection
	description Node
}

func NewExample(name string, content Node, options ...ExampleOption) Example {
	result := Example{
		PageSection: PageSection{
			name:    name,
			content: content,
		},
	}

	for _, option := range options {
		option(&result)
	}

	return result
}

func (e Example) Content() Node {
	buf := new(bytes.Buffer)
	err := Nodes(e.content).Render(buf)

	if err != nil {
		log.Fatal(err)
	}

	source := gohtml.Format(buf.String())
	source = strings.ReplaceAll(source, "&#39;", "'")

	return Article(
		Class("example"),
		ID(strcase.ToKebab(e.Name())),
		H2(Text(e.Name())),
		e.description,
		Article(
			Class("card"),
			Section(e.content),
		),
		Section(
			Class("accordion"),
			Details(
				Class("hljs source "),
				Summary(Text("HTML")),
				Pre(Code(Data("highlight", "yes"), Class("language-html"), Text(source))),
			),
		),
	)
}

func InlineCodeList(classes ...string) Node {
	result := []Node{}

	for i, c := range classes {
		if i != 0 {
			if i < len(classes)-1 {
				result = append(result, Text(", "))
			} else {
				result = append(result, Text(" and "))
			}
		}
		result = append(result, Code(Text(c)))
	}

	return Group(result)
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
