package docs

var CalloutPage = NewPage(
	"Callout",
	`<p>Callouts draw attention to supplementary information, feedback, and messages that are relevant to the current context.</p>`,
	NewExample(
		"Types",
		`<p>Use the <code>info</code>, <code>success</code>, <code>warning</code>, or <code>error</code> class to communicate the kind of message being shown.</p>`,
		`
    <aside class="callout info content">
      <strong>Info</strong>
      <p>This is some additional information.</p>
    </aside>
    <aside class="callout success content">
      <strong>Success</strong>
      <p>Your changes have been saved.</p>
    </aside>
    <aside class="callout warning content">
      <strong>Warning</strong>
      <p>This action may have unexpected consequences.</p>
    </aside>
    <aside class="callout error content">
      <strong>Error</strong>
      <p>Something went wrong. Please try again.</p>
    </aside>
    `,
	),
)
