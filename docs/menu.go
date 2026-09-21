package docs

var MenuPage = NewPage(
	"Menu",
	`<p>A menu is a popover containing a list of links or actions that a user can open from a trigger, letting them navigate or invoke actions without leaving the current context.</p>`,
	NewExample(
		"Links",
		`<p>Menu items containing an anchor can be used to navigate to other parts of the application.</p>`,
		` 
    <button class="outline" popovertarget="popover-links">Go to page</button>
    <ul id="popover-links" class="menu" popover>
	      <li><a href="/docs/about">About</a></li>
	      <li><a href="/docs/typography">Typography</a></li>
	      <li><a href="/docs/dialog">Dialog</a></li>
    </ul>
    `,
	).WithClass("grid").WithModules("Button", "Menu", "Anchor"),
	NewExample(
		"Buttons",
		`<p>Menu items containing a button can be used invoke custom actions.</p>`,
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
	).WithClass("grid").WithModules("Button", "Dialog", "Menu"),
	NewExample(
		"Checked items",
		`<p>Use <code>role="menuitemradio"</code> and <code>aria-checked</code> when a menu contains a set of mutually exclusive choices.</p>`,
		`
    <button class="outline chevron" popovertarget="popover-theme">Theme</button>
    <ul id="popover-theme" class="menu" popover role="menu" aria-label="Color theme">
      <li><button role="menuitemradio" aria-checked="true">System</button></li>
      <li><button role="menuitemradio" aria-checked="false">Light</button></li>
      <li><button role="menuitemradio" aria-checked="false">Dark</button></li>
    </ul>
    `,
	).WithClass("grid").WithModules("Button", "Menu"),
	NewExample(
		"Groups",
		`<p>You can use multiple lists to group menu items.</p>`,
		` 
    <button class="outline chevron" popovertarget="popover-groups">Go to page</button>
    <div id="popover-groups" class="menu" popover>
      <ul>
          <li><a href="/docs/about">About</a></li>
          <li><a href="/docs/modules">Modules</a></li>
      </ul>
      <ul>
          <li><a href="/docs/colors">Colors</a></li>
          <li><a href="/docs/palette">Palette</a></li>
      </ul>
      <ul>
          <li><a href="/docs/dialog">Dialog</a></li>
          <li><a href="/docs/menu">Menu</a></li>
          <li><a href="/docs/notification">Notification</a></li>
      </ul>
    </div>
    `,
	).WithClass("grid").WithModules("Button", "Menu", "Anchor"),
	NewExample(
		"Chevron",
		`<p>You can inject a chevron icon to the trigger by adding the <code>.chevron</code> class.</p>`,
		` 
    <button class="outline chevron" popovertarget="popover-trigger">Go to page</button>
    <ul id="popover-trigger" class="menu" popover>
	      <li><a href="/docs/about">About</a></li>
	      <li><a href="/docs/typography">Typography</a></li>
	      <li><a href="/docs/dialog">Dialog</a></li>
    </ul>
    `,
	).WithClass("grid").WithModules("Button", "Menu", "Anchor"),
)
