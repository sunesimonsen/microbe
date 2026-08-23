package docs

var PaginationPage = NewPage(
	"Pagination",
	`<p>Pagination lets users move through a collection of results. Use cursor-based pagination when the total number of pages is unknown, or indexed pagination when users need to jump directly to a known page.</p>`,
	NewExample(
		"Cursor-based",
		`<p>Cursor-based pagination uses an opaque cursor to request the next or previous set of results. Since the total number of pages is not needed, provide only the navigation links that are available.</p>`,
		`
    <nav class="pagination" aria-label="Results pages">
      <a class="first text" href="#">First</a>
      <a class="previous text" href="#">Previous</a>
      <a class="next text" href="#">Next</a>
      <a class="last text" href="#">Last</a>
    </nav>
    `,
	),
	NewExample(
		"Indexed",
		`<p>Indexed pagination presents numbered links for collections with a known number of pages. Use <code>aria-current="page"</code> to identify the page being viewed.</p>`,
		`
    <nav class="pagination" aria-label="Results pages">
      <a class="previous" title="Previous page" href="#"></a>
      <a href="#">1</a>
      <a href="#">2</a>
      <a href="#" aria-current="page">3</a>
      <a href="#">4</a>
      <a href="#">5</a>
      <a class="next" title="Next page" href="#"></a>
    </nav>
    `,
	),
	NewExample(
		"Ellipsis",
		`<p>Indicate gaps in the range using ellipsis.</p>`,
		`
    <nav class="pagination" aria-label="Results pages">
      <a class="first" title="First page" href="#"></a>
      <a class="previous" title="Previous page" href="#"></a>
      <a href="#">1</a>
      <span aria-label="Ellipsis indicating non-visible pages">…</span>
      <a href="#" aria-current="page">10</a>
      <span aria-label="Ellipsis indicating non-visible pages">…</span>
      <a href="#">50</a>
      <a class="next" title="Next page" href="#" ></a>
      <a class="last" title="Last page" href="#" ></a>
    </nav>
    `,
	),
)
