package docs

import (
	"errors"
	"log"
	"net/url"
	"regexp"
	"slices"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/sunesimonsen/microbe/icons"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var ErrNotFound = errors.New("Resource was not found")

type PageSection interface {
	GetNode(u url.URL) Node
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

var playgroundCSS = `
@import "https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@HEAD/assets/microbe.css";
@import "https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@HEAD/assets/microbe-accordion.css";
@import "https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@HEAD/assets/microbe-card.css";
@import "https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@HEAD/assets/microbe-dialog.css";
@import "https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@HEAD/assets/microbe-navlist.css";
@import "https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@HEAD/assets/microbe-tabs.css";
@import "https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@HEAD/assets/microbe-tag.css";
@import "https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@HEAD/assets/microbe-menu.css";

body {
  color: hsl(var(--neutral-hue) var(--neutral-saturation) var(--foreground-lightness));
  background: hsl(var(--neutral-hue) var(--neutral-saturation) var(--background-lightness));
  padding: var(--scale-5);
}

.grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(30%, 1fr));
  gap: var(--scale-4);
  justify-items: center;
  align-items: center;

  &>* {
    margin: 0;
  }
}

.grid.stretch {
  justify-items: stretch;
}

.rows {
  display: grid;
  grid-template-columns: 1fr;
  gap: var(--scale-4);
  justify-items: stretch;
}

.rows:has(.grid) {
  gap: var(--scale-6);
}

@media (max-width: 600px) {
  .grid:not(.small) {
    grid-template-columns: 1fr;
  }
}
`

func (e Example) GetNode(u url.URL) Node {
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
					Class("ghost icon jsfiddle"),
					icons.JSFiddleIcon(),
				),
				Button(
					Class("ghost icon show-source"),
					Aria("pressed", "false"),
					icons.CodeSlashIcon(),
				),
			),
		),
		Form(
			Action("https://jsfiddle.net/api/post/library/pure/"),
			Method("post"),
			Target("_blank"),
			Attr("hidden"),
			Class("jsfiddle"),
			Input(Type("hidden"), Name("title"), Value("Microbe "+u.Path+"#"+e.Name)),
			Input(Type("hidden"), Name("description"), Value("See https://microbe.sune.one for more information")),
			Input(Type("hidden"), Name("html"), Value("")),
			Input(Type("hidden"), Name("js"), Value("")),
			Input(Type("hidden"), Name("css"), Value(strings.TrimSpace(playgroundCSS))),
		),
	)
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

func (e Example) WithDescription(description ...Node) Example {
	e.Description = Group(description)
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

type GetNodeFunc func(u url.URL) Node

type CustomPageSection struct {
	Namespace string
	Name      string
	getNode   GetNodeFunc
}

func (s CustomPageSection) GetName() string {
	return s.Name
}

func (s CustomPageSection) GetNode(u url.URL) Node {
	return Section(ID(strcase.ToKebab(s.GetName())), s.getNode(u))
}

func NewPageSection(name string, getNode func(u url.URL) Node) CustomPageSection {
	return CustomPageSection{
		Name:    name,
		getNode: getNode,
	}
}

func NewStaticPageSection(name string, nodes ...Node) CustomPageSection {
	return CustomPageSection{
		Name: name,
		getNode: func(_ url.URL) Node {
			return Group(nodes)
		},
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

func (p Page) GetNode(u url.URL) Node {
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
				return s.GetNode(u)
			}),
		),
		Aside(
			Class("toc"),
			Nav(
				Class("navlist"),
				Details(
					Open(),
					Summary(Text("On this page")),
					Name("toc"),
					Ul(
						Map(p.Content, func(s PageSection) Node {
							return Li(A(Href("#"+strcase.ToKebab(s.GetName())), Text(s.GetName())))
						}),
					),
				),
			),
		),
	}
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

func (cs Categories) FindPage(name string) (Page, error) {
	for _, c := range cs {
		if page, err := c.FindPage(name); err == nil {
			return page, nil
		}
	}

	return Page{}, ErrNotFound
}

func PageHref(c Category, p Page) string {
	return "/docs/" + strcase.ToKebab(p.Name)
}

func (cs Categories) Filter(query string) Categories {
	result := Categories{}

	terms := strings.Split(query, " ")

	for _, c := range cs {
		fc := Category{
			Name: c.Name,
		}

		for _, p := range c.Pages {
			var hasMatch bool

			for _, term := range terms {
				t := strings.ToLower(term)

				if strings.Contains(strings.ToLower(c.Name), t) ||
					strings.Contains(strings.ToLower(p.Name), t) {
					hasMatch = true
				}
			}

			if hasMatch {
				fc.Pages = append(fc.Pages, p)
			}
		}

		if len(fc.Pages) > 0 {
			result = append(result, fc)
		}
	}

	return result
}

func (cs Categories) GetMenu(currentPath string, expandAll bool) Node {
	return Map(cs, func(c Category) Node {
		open := expandAll || slices.ContainsFunc(c.Pages, func(p Page) bool {
			return PageHref(c, p) == currentPath
		})

		return Details(
			If(open, Open()),
			If(open, Data("open", "true")),
			Summary(If(expandAll, TabIndex("-1")), Text(c.Name)),
			Ul(
				Map(c.Pages, func(p Page) Node {
					href := PageHref(c, p)

					return Li(A(
						If(href == currentPath, Aria("current", "page")),
						Href(href),
						Text(p.Name),
					))
				}),
			),
		)
	})
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

func ExternalLink(href, text string) Node {
	return A(Href(href), Target("_blank"), Text(text))
}
