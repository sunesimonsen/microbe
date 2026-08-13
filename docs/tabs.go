package docs

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var TabsPage = NewPage(
	"Tabs",
	NewExample(
		"Styles",
		`
      <div role="tablist" aria-label="Select your operating system">
        <button
          role="tab"
          aria-selected="true"
          aria-controls="tab-panel"
          id="tab-1">
          Windows
        </button>
        <button
          role="tab"
          aria-selected="false"
          aria-controls="tab-panel"
          id="tab-2">
          macOS
        </button>
        <button
          role="tab"
          aria-selected="false"
          aria-controls="tab-panel"
          id="tab-3">
          Linux
        </button>
      </div>
      <section 
        role="tabpanel" 
        aria-labelledby="tab-1"
        id="tab-panel">
        <p>How to run this application on Windows</p>
      </section>
    `,
	).WithDescription(
		P(
			Text("Microbe provides styles for an "),
			ExternalLink("https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/Reference/Roles/tab_role", "ARIA tabs structure"),
			Text(", use your favorite JavaScript library to make it interactive."),
		),
	),
	NewExample(
		"Disabled",
		`
      <div role="tablist" aria-label="Select your operating system">
        <button
          role="tab"
          aria-selected="false"
          aria-controls="tab-panel"
          disabled
          id="tab-1">
          Windows
        </button>
        <button
          role="tab"
          aria-selected="false"
          aria-controls="tab-panel"
          id="tab-2">
          macOS
        </button>
        <button
          role="tab"
          aria-selected="true"
          aria-controls="tab-panel"
          id="tab-3">
          Linux
        </button>
      </div>
      <section 
        role="tabpanel" 
        aria-labelledby="tab-3"
        id="tab-panel">
        <p>How to run this application on Linux</p>
      </section>
    `,
	).WithDescription(
		P(Text("You can mark individual tabs as disabled using the "), Code(Text("disabled")), Text(" attribute")),
	),
).WithDescription(
	P(Text("Tabs let users switch between related views or sections of content within the same context, without needing to navigate to a different page.")),
)
