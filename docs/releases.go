package docs

import (
	"net/url"
	"slices"

	"github.com/iancoleman/strcase"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type Release struct {
	Version string
	Modules []string
}

var releases = []Release{
	{Version: "HEAD", Modules: []string{"Accordion", "Breadcrumb", "Card", "Dialog", "Menu", "Navlist", "Pagination", "Tabs", "Tag"}},
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
	`<p>Lets you pick a Microbe release and the optional modules to include, then generates the corresponding HTML link tags to add to your page.</p>`,
	NewPageSection(
		"Version picker",
		func(u url.URL) Node {
			currentVersion := CurrentVersion(u)
			release := GetRelease(currentVersion)

			return Group([]Node{
				Article(
					Class("card raised"),
					Header(Text("Version picker")),
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
					Header(Text("HTML snippet")),
					Section(
						Class("source"),
						Pre(
							Code(
								Class("language-html hljs language-xml"),
								Data("highlight", "yes"),
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
							),
						),
					),
					Footer(
						Class("actions"),
						Button(
							Class("ghost icon copy-source"),
							Title("Copy snippet"),
							Raw(`
                <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" fill="currentColor" class="bi bi-clipboard" viewBox="0 0 16 16">
                  <path d="M4 1.5H3a2 2 0 0 0-2 2V14a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V3.5a2 2 0 0 0-2-2h-1v1h1a1 1 0 0 1 1 1V14a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V3.5a1 1 0 0 1 1-1h1z"/>
                  <path d="M9.5 1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 .5-.5zm-3-1A1.5 1.5 0 0 0 5 1.5v1A1.5 1.5 0 0 0 6.5 4h3A1.5 1.5 0 0 0 11 2.5v-1A1.5 1.5 0 0 0 9.5 0z"/>
                </svg>
                `,
							),
						),
					),
				),
			})
		},
	),
)
