package docs

var BreadcrumbPage = NewPage(
	"Breadcrumb",
	`<p>Breadcrumbs show the current location within a hierarchy of pages, helping users understand where they are and move back to an earlier level.</p>`,
	NewExample(
		"Example",
		`<p>Use a semantic <code>nav</code> element with an accessible label and an ordered list of links. Mark the current page with <code>aria-current="page"</code>.</p>`,
		`
    <nav class="breadcrumb" aria-label="Breadcrumb">
      <ol>
        <li><a href="#">Home</a></li>
        <li><a href="#">Documentation</a></li>
        <li aria-current="page">Breadcrumb</li>
      </ol>
    </nav>
    `,
	),
)
