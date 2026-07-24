package docs

import (
	"net/url"
	"slices"

	"github.com/iancoleman/strcase"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var About = NewPage(
	"About",
	NewStaticPageSection(
		"Usage",
		H2(Text("Usage")),
		P(Text("This project is currently under construction, so no releases has been made yet. But you can find the development stylesheets here if you want to play around with it")),
		P(Text("The main stylesheet is "), Code(Text("microbe.css")),
			Text(" is required and provides base styles for most HTML elements.")),
		Pre(Code(Text("https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@HEAD/assets/microbe.css"))),
		P(Text("A number of additional stylesheets exists extending the main stylesheet with extra components. You can include these to meet your needs:")),
		Ul(
			Li(A(Href("/components/accordion"), Text("Accordion"))),
			Li(A(Href("/components/card"), Text("Card"))),
			Li(A(Href("/navigation/navlist"), Text("Navlist"))),
		),
	),
).WithDescription(
	P(Text("Microbe is a CSS framework with a core stylesheet that elegantly styles most native HTML elements, plus optional modules for common user interface styles.")),
	P(Text("The design system is built around the "), ExternalLink("https://en.wikipedia.org/wiki/Golden_ratio", "golden ratio"), Text(" and the "), ExternalLink("https://developer.mozilla.org/en-US/docs/Learn_web_development/Core/Styling_basics/Values_and_units#relative_length_units", "ralative units"), Text(" allowing styles to scale beautifully with the surrounding font size.")),
	P(Text("It ships with opinionated defaults, but you can easily override them using "), ExternalLink("https://developer.mozilla.org/en-US/docs/Web/CSS/Guides/Cascading_variables/Using_custom_properties", "CSS variables"), Text(" and "), ExternalLink("https://developer.mozilla.org/en-US/docs/Web/CSS/Reference/At-rules/@layer", "CSS layers"), Text(".")),
)

type Release struct {
	Version string
	Modules []string
}

var releases = []Release{
	{Version: "HEAD", Modules: []string{"Accordion", "Card", "Navlist"}},
}

func CurrentVersion(u url.URL) string {
	currentVersion := u.Query().Get("version")

	if currentVersion == "" {
		return "HEAD"
	}

	return currentVersion
}

func CurrentModules(u url.URL) []string {
	q := u.Query()
	modules := q["modules"]
	return modules
}

func IncludesModule(u url.URL, name string) bool {
	modules := CurrentModules(u)

	return slices.Contains(modules, name)
}

func GetRelease(version string) Release {
	releaseIndex := slices.IndexFunc(releases, func(r Release) bool {
		return r.Version == version
	})

	release := releases[0]
	if releaseIndex != -1 {
		release = releases[releaseIndex]
	}

	return release
}

var ReleasesPage = NewPage(
	"Releases",
	NewPageSection(
		"Version picker",
		func(u url.URL) Node {
			currentVersion := CurrentVersion(u)
			release := GetRelease(currentVersion)

			return Group([]Node{
				H2(Text("Version picker")),
				Article(
					Class("card raised"),
					Section(
						Form(
							Label(
								Text("Version"),
								Select(
									Name("version"),
									Attr("onchange", "this.form.submit()"),
									Map(releases, func(r Release) Node {
										return Option(
											Value(r.Version),
											Text(r.Version),
											If(currentVersion == r.Version, Selected()),
										)
									}),
								),
							),
							FieldSet(
								Legend(Text("Include modules")),
								Map(release.Modules, func(module string) Node {
									id := strcase.ToKebab(module)
									return Label(
										Input(Type("checkbox"),
											Name("modules"),
											Attr("onchange", "this.form.submit()"),
											Value(id),
											If(IncludesModule(u, id), Checked()),
										),
										Text(module),
									)
								}),
							),
						),
					),
				),
			})
		},
	),
	NewPageSection(
		"HTML snippet",
		func(u url.URL) Node {
			currentVersion := CurrentVersion(u)
			release := GetRelease(currentVersion)

			return Group([]Node{
				H3(Text("HTML snippet")),
				Pre(Code(
					Textf("<link href=\"https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@%s/assets/microbe.css\" rel=\"stylesheet\" type=\"text/css\">\n",
						currentVersion,
					),
					Map(release.Modules, func(module string) Node {
						id := strcase.ToKebab(module)
						if !IncludesModule(u, id) {
							return nil
						}

						return Textf(
							"<link href=\"https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@%s/assets/microbe-%s.css\" rel=\"stylesheet\" type=\"text/css\">\n",
							currentVersion,
							id,
						)
					}),
				)),
			})
		},
	),
)
