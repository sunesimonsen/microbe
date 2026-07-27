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
		P(Text("Pick a release on the "), A(Href("/getting-started/releases"), Text("release page")), Text(" and just start adding HTML markup described by the examples.")),
	),
	NewExample(
		"Example",
		`
    <form id="getting-started-demo">
      <fieldset>
        <label>
          Name
          <input name="name" placeholder="Name" autocomplete="name">
        </label>
        <label>
          Email
          <input name="email" placeholder="Email" autocomplete="email" aria-describedby="email-hint">
          <small id="email-hint">
            We’ll never share your email with anyone else.
          </small>
        </label>
        <label>
          <input type="checkbox" name="newsletter" aria-describedby="newsletter-hint" checked>
          Newsletter
        </label>
        <small id="newsletter-hint">
          We will send you a newsletter every week
        </small>
      </fieldset>
      <div class="actions">
        <button class="outline" type="reset">Reset</button>
        <button class="solid" type="submit">Submit</button>
      </div>
    </form>
    <script>
      let index = 0
      setInterval(() => {
        const form = document.getElementById("getting-started-demo")

        index = (index + 1) % 4;
        form.classList.value = "overrides-" + index;
      }, 3000);
    </script>
    <style>
    @layer overrides {
      #getting-started-demo {
        & * {
          transition: all ease-in-out 1s;
        }

        &.overrides-1,
        &.overrides-2 {
          --accent-hue: 170;
          --accent-saturation: 40%;
          & button,
          & input:not([type=checkbox]) {
            border-radius: var(--scale-6);
          }
        }

        &.overrides-2 {
          --accent-hue: 300;

          button.solid {
            box-shadow: 0 0 var(--scale-3) var(--scale-3) hsl(var(--accent-hue) var(--accent-saturation) var(--lightness-3));
          }
        }

        &.overrides-3,
        &.overrides-4 {
          --accent-hue: 20;

          & button,
          & input {
            border-radius: 0;
          }
        }
      }
    }
    </style>
    `,
	).WithDescription(
		P(Text("Here is a quick example showing the default styles and potential overrides.")),
		P(Text("You can change the accent color and size on all elements and customize any CSS with overrides.")),
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
	{Version: "HEAD", Modules: []string{"Accordion", "Card", "Navlist", "Tabs"}},
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
                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-clipboard" viewBox="0 0 16 16">
                  <path d="M4 1.5H3a2 2 0 0 0-2 2V14a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V3.5a2 2 0 0 0-2-2h-1v1h1a1 1 0 0 1 1 1V14a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V3.5a1 1 0 0 1 1-1h1z"/>
                  <path d="M9.5 1a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 .5-.5zm-3-1A1.5 1.5 0 0 0 5 1.5v1A1.5 1.5 0 0 0 6.5 4h3A1.5 1.5 0 0 0 11 2.5v-1A1.5 1.5 0 0 0 9.5 0z"/>
                </svg>
              `),
						),
					),
				),
			})
		},
	),
)
