# Microbe — project map

## Purpose

Microbe is an in-progress Go web application that documents and serves **Microbe.css**, an opinionated CSS library based on native HTML, relative `em` sizing, golden-ratio spacing, CSS variables, and CSS layers. The documentation examples are authored in Go and rendered server-side.

## Stack and runtime

- Go `1.26`; module: `github.com/sunesimonsen/microbe`
- HTTP: `net/http` + `github.com/go-chi/chi/v5`
- HTML rendering: `maragu.dev/gomponents`
- Name-to-URL normalization: `github.com/iancoleman/strcase`
- Browser enhancement: vanilla JavaScript + HTMX `2.0.10` loaded from jsDelivr
- Default server: `localhost:8080`; static files are served from the relative `assets/` directory
- Production image: multi-stage Alpine Docker build; executable is `/app/main`

## Repository layout

```text
.
├── main.go                 # process entry point; creates server and listens on :8080
├── server/                 # chi router, HTTP handlers, route tests
│   ├── server.go           # Server wrapper and NewServer
│   ├── routes.go           # middleware and route registration
│   ├── handlers.go         # docs/search rendering and error handling
│   └── routes_test.go      # endpoint behavior tests
├── docs/                   # documentation model, registry, pages, examples
│   ├── structure.go        # Page/Section/Example/Category types and menu/search logic
│   ├── index.go            # ordered documentation registry (docs.Index)
│   ├── structure_test.go   # category filtering tests
│   └── <page>.go            # one Go source file per documentation page
├── views/                  # shared page/layout components built with gomponents
│   ├── page.go             # HTML document shell, metadata, styles/scripts
│   └── layout.go           # header, search dialog, docs menu, result layouts
├── icons/                  # small inline SVG gomponents (Burger, CodeSlash, Github,
│                           # JSFiddle, Search)
├── assets/                 # CSS library, documentation CSS/JS, favicons and manifest
├── Makefile                # build, run, test, dev, coverage, deploy, browse tasks
├── Dockerfile              # Go build stage + minimal Alpine runtime image
├── go.mod / go.sum         # dependencies and checksums
├── README.md               # minimal project description
└── .air.toml / .editorconfig # live reload and formatting/editor settings
```

`tmp/`, `.pi-tasks/`, local logs, and build outputs are workspace artifacts rather than application source. `tmp/main` is the Makefile build target and is not required in source control.

## Request flow

1. `main.go` calls `server.NewServer()` and passes the returned `Server` to `http.ListenAndServe(":8080", ...)`.
2. `server/routes.go` creates a chi router with slash redirect and request logging middleware.
3. `/` and `/getting-started/about` use `IndexHandler`; `/` permanently redirects to `/docs/about`.
4. `/docs/{page}` uses `DocsHandler`, which finds a page in `docs.Index`, renders `views.DocsLayout`, and returns 404 for `docs.ErrNotFound`.
5. `/search?query=...` filters `docs.Index` by category/page names (case-insensitive, any term may match), then returns a partial navlist for HTMX. The `Referer` path controls the current/expanded menu state.
6. `/assets/*` serves files from `assets/` after stripping the URL prefix.
7. `views.Page` emits the full HTML shell and loads the core stylesheet, all current optional module stylesheets, `page.js`, and external HTMX. `docs.Page.GetNode` adds Highlight.js assets for code examples.

## Documentation architecture

`docs.Index` is an ordered `Categories` value. Add a page in two steps: define a `Page` in `docs/<topic>.go`, then include it in the appropriate `NewCategory(...)` in `docs/index.go`.

```text
Getting started: About, Releases
Theming:         Colors, Palette
Layout:          Accordion, Card, Tabs, Spacing
Content:         Typography, List, Table
Navigation:      Anchor, Navlist
Forms:           Button, Checkbox, Input, Textarea, Radio, Range, Select, Switch
Loaders:         Progress
Popups:          Dialog, Menu
Data:            Tag
```

Page URLs are generated as `/docs/<kebab-case page name>`; lookup is case/format tolerant through `strcase.ToKebab`. A `Page` contains `PageSection`s. Sections are either:

- `NewExample(name, descriptionHTML, sourceHTML)`, optionally decorated with `.WithClass(...)`; source is whitespace-normalized and can be shown/highlighted or posted to JSFiddle.
- `NewStaticPageSection(...)` / `NewPageSection(...)` for static or URL-aware custom content.

`Page.GetNode` renders the title/description, all sections, and a table of contents. `Categories.GetMenu` creates the collapsible navigation; `Categories.Filter` searches only category and page names (not page descriptions/example source). Releases currently exposes only `HEAD`, with optional module names defined in `docs/releases.go`.

## Frontend/assets

- `assets/microbe.css`: core CSS reset/defaults, typography, forms, colors, spacing, buttons, tables, progress, etc.; organized with `@layer settings, defaults, modules, overrides`.
- `assets/microbe-{accordion,card,dialog,menu,navlist,tabs,tag}.css`: opt-in component modules.
- `assets/page.css`: documentation-site layout and example/playground styles.
- `assets/page.js`: color playground controls, indeterminate checkbox setup, Highlight.js activation, source toggle, JSFiddle export, clipboard actions, and color-sample copying.
- Remaining assets are favicon/PWA manifest files.

Examples intentionally use semantic/native HTML and accessibility attributes (`fieldset`/`legend`, labels, ARIA tabs, `aria-current`, `aria-describedby`, validation state, native dialog/popover commands); Microbe.css supplies styling while JavaScript supplies only site/demo behavior.

## Common commands

```sh
make test       # go test ./...
make build      # go build -o tmp/main .
make run        # go run .
make dev        # Test=true air (live reload)
make cover      # coverage report and cover.html
make clean      # remove tmp/
go test ./...    # direct equivalent of make test
```

`make deploy` pushes `main` to the configured `dokku` remote; `make browse` opens the hosted documentation site. Run the server from the repository root so the relative `assets/` path resolves.

## Change guide

- Documentation content belongs in `docs/`; keep the registry order in `docs/index.go` synchronized with the desired navigation order.
- Shared HTML belongs in `views/`; request/routing behavior belongs in `server/`; reusable inline SVG belongs in `icons/`.
- Core CSS changes go in `assets/microbe.css`; component-specific CSS goes in its corresponding `microbe-*.css`; documentation-only styling goes in `assets/page.css`.
- Keep examples as raw HTML strings inside page definitions so the rendered demo and displayed source stay aligned.
- Run `make test` after changes. Route tests assume the process is run from the repository root and verify redirects, known/missing docs, search, malformed referers, and trailing-slash handling.
