package docs

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var MenuPage = NewPage(
	"Menu",
	NewExample(
		"Links",
		` 
    <button class="outline" popovertarget="popover-links">Go to page</button>
    <ul id="popover-links" class="menu" popover>
	      <li><a href="/docs/about">About</a></li>
	      <li><a href="/docs/typography">Typography</a></li>
	      <li><a href="/docs/dialog">Dialog</a></li>
    </ul>
    `,
	).WithDescription(
		P(Text("Menu items containing an anchor can be used to navigate to other parts of the application.")),
	),
	NewExample(
		"Buttons",
		`
    <button class="outline" popovertarget="popover-actions">Actions</button>
    <ul id="popover-actions" class="menu" popover>
      <li><button autofocus commandfor="example-dialog" command="show-modal">Open dialog</button></li>
      <li><button onclick="javascript:alert('You clicked me')">Custom JavaScript</button></li>
      <li><button onclick="javascript:alert('You clicked me')" popovertarget="popover-actions" popovercommand="close">Custom JavaScript and closing</button></li>
      <li><button popovertarget="popover-actions" popovercommand="close">Close menu</button></li>
    </ul>
    <dialog id="example-dialog" class="small" closedby="any">
      <header>Example</header>
      <section>I was opened by a menu action.</section>
      <footer class="actions">
        <button class="solid" command="close" commandfor="example-dialog">
          Close
        </button>
      </footer>
    </dialog>
    `,
	).WithDescription(
		P(Text("Menu items containing a button can be used invoke custom actions.")),
	),
	NewExample(
		"Chevron",
		` 
    <button class="outline chevron" popovertarget="popover-trigger">Go to page</button>
    <ul id="popover-trigger" class="menu" popover>
	      <li><a href="/docs/about">About</a></li>
	      <li><a href="/docs/typography">Typography</a></li>
	      <li><a href="/docs/dialog">Dialog</a></li>
    </ul>
    `,
	).WithDescription(
		P(Text("You can inject a chevron icon to the trigger by adding the "), InlineCodeList(".chevron"), Text(" class.")),
	),
).WithDescription(
	P(Text("A menu is a popover containing a list of links or actions that a user can open from a trigger, letting them navigate or invoke actions without leaving the current context.")),
)
