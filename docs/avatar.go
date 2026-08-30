package docs

var AvatarPage = NewPage(
	"Avatar",
	`<p>Avatars represent people, teams, or organizations using an image, icon, or initials.</p>`,
	NewExample(
		"Shape",
		`<p>Use the default <code>avatar</code> class for people and other individuals. It renders as a circle. Add the <code>logo</code> class for organizations or brands when a square shape is more appropriate.</p>`,
		`
    <figure class="avatar" aria-label="User avatar">
      <svg aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" fill="currentColor" class="bi bi-person" viewBox="0 0 16 16">
        <path d="M8 8a3 3 0 1 0 0-6 3 3 0 0 0 0 6m2-3a2 2 0 1 1-4 0 2 2 0 0 1 4 0m4 8c0 1-1 1-1 1H3s-1 0-1-1 1-4 6-4 6 3 6 4m-1-.004c-.001-.246-.154-.986-.832-1.664C11.516 10.68 10.289 10 8 10s-3.516.68-4.168 1.332c-.678.678-.83 1.418-.832 1.664z"/>
      </svg>
    </figure>
    <figure class="avatar logo" aria-label="Google logo">
      <svg aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" fill="currentColor" class="bi bi-google" viewBox="0 0 16 16">
        <path d="M15.545 6.558a9.4 9.4 0 0 1 .139 1.626c0 2.434-.87 4.492-2.384 5.885h.002C11.978 15.292 10.158 16 8 16A8 8 0 1 1 8 0a7.7 7.7 0 0 1 5.352 2.082l-2.284 2.284A4.35 4.35 0 0 0 8 3.166c-2.087 0-3.86 1.408-4.492 3.304a4.8 4.8 0 0 0 0 3.063h.003c.635 1.893 2.405 3.301 4.492 3.301 1.078 0 2.004-.276 2.722-.764h-.003a3.7 3.7 0 0 0 1.599-2.431H8v-3.08z"/>
      </svg>
    </figure>
    `,
	).WithClass("grid small"),
	NewExample(
		"Type",
		`<p>An avatar can contain an icon, an image, or a short text label such as a user's initials. Give avatars an accessible name, and provide meaningful alternative text when using an image.</p>`,
		`
    <figure class="avatar" aria-label="Avatar for Sune Simonsen">
      <svg aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" fill="currentColor" class="bi bi-person" viewBox="0 0 16 16">
        <path d="M8 8a3 3 0 1 0 0-6 3 3 0 0 0 0 6m2-3a2 2 0 1 1-4 0 2 2 0 0 1 4 0m4 8c0 1-1 1-1 1H3s-1 0-1-1 1-4 6-4 6 3 6 4m-1-.004c-.001-.246-.154-.986-.832-1.664C11.516 10.68 10.289 10 8 10s-3.516.68-4.168 1.332c-.678.678-.83 1.418-.832 1.664z"/>
      </svg>
    </figure>
    <img class="avatar" src="https://microbe.sune.one/assets/nyancat.jpg" alt="Avatar for Sune Simonsen">
    <span class="avatar" aria-label="Avatar for Sune Simonsen">SSS</span>
    `,
	).WithClass("grid small"),
	NewExample(
		"Size",
		`<p>As any element in Microbe you can easily change it's size by setting the font-size of the element.</p>`,
		`
    <div class="grid small">
      <figure class="avatar" aria-label="Avatar for Sune Simonsen" style="font-size: smaller">
        <svg aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" fill="currentColor" class="bi bi-person" viewBox="0 0 16 16">
          <path d="M8 8a3 3 0 1 0 0-6 3 3 0 0 0 0 6m2-3a2 2 0 1 1-4 0 2 2 0 0 1 4 0m4 8c0 1-1 1-1 1H3s-1 0-1-1 1-4 6-4 6 3 6 4m-1-.004c-.001-.246-.154-.986-.832-1.664C11.516 10.68 10.289 10 8 10s-3.516.68-4.168 1.332c-.678.678-.83 1.418-.832 1.664z"/>
        </svg>
      </figure>
      <figure class="avatar" aria-label="Avatar for Sune Simonsen">
        <svg aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" fill="currentColor" class="bi bi-person" viewBox="0 0 16 16">
          <path d="M8 8a3 3 0 1 0 0-6 3 3 0 0 0 0 6m2-3a2 2 0 1 1-4 0 2 2 0 0 1 4 0m4 8c0 1-1 1-1 1H3s-1 0-1-1 1-4 6-4 6 3 6 4m-1-.004c-.001-.246-.154-.986-.832-1.664C11.516 10.68 10.289 10 8 10s-3.516.68-4.168 1.332c-.678.678-.83 1.418-.832 1.664z"/>
        </svg>
      </figure>
      <figure class="avatar" aria-label="Avatar for Sune Simonsen" style="font-size: larger">
        <svg aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" fill="currentColor" class="bi bi-person" viewBox="0 0 16 16">
          <path d="M8 8a3 3 0 1 0 0-6 3 3 0 0 0 0 6m2-3a2 2 0 1 1-4 0 2 2 0 0 1 4 0m4 8c0 1-1 1-1 1H3s-1 0-1-1 1-4 6-4 6 3 6 4m-1-.004c-.001-.246-.154-.986-.832-1.664C11.516 10.68 10.289 10 8 10s-3.516.68-4.168 1.332c-.678.678-.83 1.418-.832 1.664z"/>
        </svg>
      </figure>
    </div>
    <div class="grid small">
      <img class="avatar" src="https://microbe.sune.one/assets/nyancat.jpg" alt="Avatar for Sune Simonsen" style="font-size: smaller">
      <img class="avatar" src="https://microbe.sune.one/assets/nyancat.jpg" alt="Avatar for Sune Simonsen">
      <img class="avatar" src="https://microbe.sune.one/assets/nyancat.jpg" alt="Avatar for Sune Simonsen" style="font-size: larger">
    </div>
    <div class="grid small">
      <span class="avatar" aria-label="Avatar for Sune Simonsen" style="font-size: smaller">SSS</span>
      <span class="avatar" aria-label="Avatar for Sune Simonsen">SSS</span>
      <span class="avatar" aria-label="Avatar for Sune Simonsen" style="font-size: larger">SSS</span>
    </div>
    `,
	).WithClass("rows"),
	NewExample(
		"Interactive",
		`<p>Anchors and buttons are allowed to be styled as avatars, this is useful for opening a user profile or linking to a external pages.</p>`,
		`
    <button class="avatar" aria-label="User avatar" command="show-modal" commandfor="example-profile">
      <svg aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" fill="currentColor" class="bi bi-person" viewBox="0 0 16 16">
        <path d="M8 8a3 3 0 1 0 0-6 3 3 0 0 0 0 6m2-3a2 2 0 1 1-4 0 2 2 0 0 1 4 0m4 8c0 1-1 1-1 1H3s-1 0-1-1 1-4 6-4 6 3 6 4m-1-.004c-.001-.246-.154-.986-.832-1.664C11.516 10.68 10.289 10 8 10s-3.516.68-4.168 1.332c-.678.678-.83 1.418-.832 1.664z"/>
      </svg>
    </button>
    <a class="avatar logo" aria-label="Google logo" href="https://www.google.com" target="_blank">
      <svg aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" fill="currentColor" class="bi bi-google" viewBox="0 0 16 16">
        <path d="M15.545 6.558a9.4 9.4 0 0 1 .139 1.626c0 2.434-.87 4.492-2.384 5.885h.002C11.978 15.292 10.158 16 8 16A8 8 0 1 1 8 0a7.7 7.7 0 0 1 5.352 2.082l-2.284 2.284A4.35 4.35 0 0 0 8 3.166c-2.087 0-3.86 1.408-4.492 3.304a4.8 4.8 0 0 0 0 3.063h.003c.635 1.893 2.405 3.301 4.492 3.301 1.078 0 2.004-.276 2.722-.764h-.003a3.7 3.7 0 0 0 1.599-2.431H8v-3.08z"/>
      </svg>
    </a>
    <dialog id="example-profile" class="small" closedby="any">
      <header>
        Sune Simonsen
        <button rel="prev" aria-label="Close" commandfor="example-profile" command="close" tabindex="1"></button>
      </header>
      <section>
        <p>
          Vestibulum venenatis neque nec iaculis viverra. Proin a odio ex.
          Etiam eget rutrum nulla, id finibus tellus. Proin dignissim tortor
          lacus, in vehicula arcu commodo ut. Aenean quis euismod lectus, vitae
          ornare ex. Vivamus posuere lectus at quam maximus semper. Fusce
          sagittis in est a consequat.
        </p>
      </section>
    </dialog>
    `,
	).WithClass("grid small"),
)
