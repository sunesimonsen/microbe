package docs

var NotificationPage = NewPage(
	"Notification",
	`<p>Notifications provide brief, non-modal feedback that can be collected in a fixed toast container.</p>`,
	NewExample(
		"Types",
		`<p>Use the <code>info</code>, <code>success</code>, <code>warning</code>, or <code>error</code> class to communicate the kind of message being shown.</p>`,
		`
    <aside class="notification info">
      <header>
        <strong>Welcome back!</strong>
        <button class="close" aria-label="Dismiss notification"></button>
      </header>
      <section>
        <p>Your reading list is ready whenever you are.</p>
      </section>
    </aside>
    <aside class="notification success">
      <header>
        <strong>Changes saved</strong><br>
        <button class="close" aria-label="Dismiss notification"></button>
      </header>
      <section>
        <p>Your profile is up to date.</p>
      </section>
    </aside>
    <aside class="notification warning">
      <header>
        <strong>Heads up</strong><br>
        <button class="close" aria-label="Dismiss notification"></button>
      </header>
      <section>
        <p>Your trial ends in 3 days. Add a payment method to keep your workspace active.</p>
      </section>
    </aside>
    <aside class="notification error">
      <header>
        <strong>Couldn’t sync</strong>
        <button class="close" aria-label="Dismiss notification"></button>
      </header>
      <section>
        <p>We couldn’t save your latest changes. Check your connection and try again.</p>
      </section>
    </aside>
    `,
	).WithClass("rows"),
	NewExample(
		"Interactive",
		`<p>Create notifications dynamically in a toast area rendered in the popover layer. Choose their type, then dismiss them manually or let them disappear after five seconds.</p>`,
		`
    <form id="notification-controls">
      <label>
        Notification type
        <select id="notification-type">
          <option value="info">Info</option>
          <option value="success">Success</option>
          <option value="warning">Warning</option>
          <option value="error">Error</option>
        </select>
      </label>
      <div class="actions">
        <button class="solid" type="button">Add notification</button>
      </div>
    </form>
    <aside class="toast" role="status" aria-live="polite" popover="manual"></aside>
    <template id="notification-template">
      <aside class="notification">
        <header>
          <strong></strong>
          <button class="close" type="button" aria-label="Dismiss notification"></button>
        </header>
        <section>
          <p><span></span></p>
        </section>
      </aside>
    </template>
    <script>
      const form = document.getElementById("notification-controls")
      const type = document.getElementById("notification-type")
      const template = document.getElementById("notification-template")
      const playground = form.parentElement

      const toast = playground.querySelector(".toast")
      toast.showPopover()

      const messages = {
        info: "Here is some useful information.",
        success: "Your changes were saved successfully.",
        warning: "Please review this before continuing.",
        error: "Something went wrong. Please try again."
      }

      form.querySelector("button").addEventListener("click", () => {
        const notification = template.content.firstElementChild.cloneNode(true)
        notification.classList.add(type.value)
        notification.querySelector("strong").textContent = type.options[type.selectedIndex].textContent
        notification.querySelector("span").textContent = messages[type.value]
        toast.append(notification)

        setTimeout(() => {
          if (notification.isConnected) {
            notification.remove()
          }
        }, 5000)
      })

      playground.addEventListener("click", (event) => {
        const close = event.target.closest(".notification .close")

        if (close) {
          close.closest(".notification").remove()
        }
      })
    </script>
    `,
	),
)
