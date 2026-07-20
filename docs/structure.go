package docs

import (
	"errors"
	"io"
	"log"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/sunesimonsen/microbe/icons"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var ErrNotFound = errors.New("Resource was not found")

type PageSection interface {
	Node
	GetName() string
}

type Example struct {
	Name        string
	Description Node
	Class       string
	Source      string
}

func (e Example) GetName() string {
	return e.Name
}

func (e Example) Render(w io.Writer) error {
	return Article(
		Class("example"),
		ID(strcase.ToKebab(e.Name)),
		H2(Text(e.Name)),
		e.Description,
		Section(
			Class("card raised"),
			Section(If(e.Class != "", Class(e.Class)), Raw(string(e.Source))),
			Footer(
				Class("actions"),
				Button(
					Class("ghost icon show-source js"),
					Aria("pressed", "false"),
					icons.CodeSlashIcon(),
				),
			),
		),
	).Render(w)
}

func TrimCommonWhitespace(text string) string {
	lines := strings.Split(text, "\n")

	wsr, err := regexp.Compile(`^\s*`)
	if err != nil {
		log.Fatal(err)
	}

	cl := -1
	cws := ""

	for _, line := range lines {
		ws := wsr.Find([]byte(line))
		if 0 < len(ws) && len(ws) < len(line) && (len(ws) < cl || cl == -1) {
			cl = len(ws)
			cws = string(ws)
		}
	}

	if cl > 0 {
		for i, line := range lines {
			ul, found := strings.CutPrefix(line, cws)

			if found {
				lines[i] = ul
			}
		}
	}

	return strings.TrimSpace(strings.Join(lines, "\n"))
}

func (e Example) WithDescription(description Node) Example {
	e.Description = description
	return e
}

func (e Example) WithClass(class string) Example {
	e.Class = class
	return e
}

func NewExample(name string, source string) Example {
	return Example{
		Name:   name,
		Source: TrimCommonWhitespace(source),
	}
}

type CustomPageSection struct {
	Namespace string
	Name      string
	Node      Node
}

func (s CustomPageSection) GetName() string {
	return s.Name
}

func (s CustomPageSection) Render(w io.Writer) error {
	return s.Node.Render(w)
}

func NewPageSection(name string, node Node) CustomPageSection {
	return CustomPageSection{
		Name: name,
		Node: node,
	}
}

type Page struct {
	Name        string
	Description Node
	Content     []PageSection
}

func (p Page) FindSection(name string) (PageSection, error) {
	for _, s := range p.Content {
		if strcase.ToKebab(s.GetName()) == strcase.ToKebab(name) {
			return s, nil
		}
	}

	return CustomPageSection{}, ErrNotFound
}

func (p Page) Render(w io.Writer) error {
	return Group{
		Link(Rel("stylesheet"), Href("https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.11.1/styles/codepen-embed.min.css")),
		Script(Src("https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.11.1/highlight.min.js")),
		HGroup(
			H1(Text(p.Name)),
			p.Description,
		),
		Div(
			Role("document"),
			Map(p.Content, func(s PageSection) Node {
				return s
			}),
		),
		Aside(
			Class("toc"),
			Nav(
				Class("navlist"),
				Details(
					Open(),
					Summary(Text("Content")),
					Name("toc"),
					Ul(
						Map(p.Content, func(s PageSection) Node {
							return Li(A(Href("#"+strcase.ToKebab(s.GetName())), Text(s.GetName())))
						}),
					),
				),
			),
		),
	}.Render(w)
}

func (p Page) WithDescription(description ...Node) Page {
	p.Description = Group(description)
	return p
}

func NewPage(name string, content ...PageSection) Page {
	return Page{
		Name:    name,
		Content: content,
	}
}

type Category struct {
	Name  string
	Pages []Page
}

func (c Category) FindPage(name string) (Page, error) {
	for _, p := range c.Pages {
		if strcase.ToKebab(p.Name) == strcase.ToKebab(name) {
			return p, nil
		}
	}

	return Page{}, ErrNotFound
}

func NewCategory(name string, pages ...Page) Category {
	return Category{
		Name:  name,
		Pages: pages,
	}
}

type Categories []Category

func (cs Categories) FindCategory(name string) (Category, error) {
	for _, c := range cs {
		if strcase.ToKebab(c.Name) == strcase.ToKebab(name) {
			return c, nil
		}
	}

	return Category{}, ErrNotFound
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
