package docs

var PopoverPage = NewPage(
	"Popover",
	`<p>Popovers are non-modal overlays that display content in the top layer, positioned near the control that opens them. They use the native popover API and are dismissed when the user clicks outside them or presses Escape.</p>`,
	NewExample(
		"Basic",
		`<p>Connect a button to a <code>div</code> with the <code>popovertarget</code> attribute. The popover is positioned below the button when there is enough room.</p>`,
		`
    <button class="outline" popovertarget="basic-popover">Open popover</button>
    <div id="basic-popover" class="popover" popover>
      <p>This is a popover with some additional information.</p>
    </div>
    `,
	).WithModules("Button", "Popover"),
	NewExample(
		"Content",
		`<p>A popover is only styled as a container. Its contents remain unstyled, so you can use the HTML elements and components that suit your content.</p>`,
		`
    <button class="outline" popovertarget="content-popover">Account</button>
    <div id="content-popover" class="popover" popover>
      <h3>Account</h3>
      <p>You are signed in as <strong>Alex Morgan</strong>.</p>
      <a href="#">View profile</a>
    </div>
    `,
	).WithModules("Button", "Popover"),
)
