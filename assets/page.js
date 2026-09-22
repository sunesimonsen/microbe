const htmlTemplate = (version, modules, source) => `\
<link rel="stylesheet" type="text/css" href="https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@${version}/assets/microbe.css">
${modules.map((module) => `<link rel="stylesheet" type="text/css" href="https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@${version}/assets/microbe-${module}.css">`).join("\n")}

${source}`

const releaseTemplate = (version, modules) => [
  `<link href="https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@${version}/assets/microbe.css" rel="stylesheet" type="text/css">`,
  ...modules.map((module) => `<link href="https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@${version}/assets/microbe-${module}.css" rel="stylesheet" type="text/css">`),
].join("\n") + "\n"

const cssTemplate = `\
body {
  color: hsl(var(--neutral-hue) var(--neutral-saturation) var(--foreground-lightness));
  background: hsl(var(--neutral-hue) var(--neutral-saturation) var(--background-lightness));
  padding: var(--scale-5);
}

.grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(30%, 1fr));
  gap: var(--scale-4);
  justify-items: center;
  align-items: center;

  &>*:not(dialog, [popover]) {
    margin: 0;
  }
}

.grid.stretch {
  justify-items: stretch;
}

.rows {
  display: grid;
  grid-template-columns: 1fr;
  gap: var(--scale-4);
  justify-items: stretch;
}

.rows:has(.grid) {
  gap: var(--scale-6);
}

@media (max-width: 600px) {
  .grid:not(.small) {
    grid-template-columns: 1fr;
  }
}
`

const indent = (source, spacing) => source.replace(/^/gm, spacing)

const wrapWithClasses = (source, classes) => `<section class="${classes.join(" ")}">
${indent(source, '  ')}
</section>`

if (window.self === window.top) {
  window.addEventListener('load', () => {
    const colorSchemeMenu = document.getElementById("color-scheme-menu")
    if (colorSchemeMenu) {
      const colorSchemeStorageKey = "microbe-color-scheme"
      const colorSchemeButtons = colorSchemeMenu.querySelectorAll("[data-color-scheme]")
      let selectedColorScheme = "system"

      try {
        const storedColorScheme = localStorage.getItem(colorSchemeStorageKey)
        if (["system", "light", "dark"].includes(storedColorScheme)) {
          selectedColorScheme = storedColorScheme
        }
      } catch (_) {
        // Local storage may be unavailable in privacy-restricted browsers.
      }

      const applyColorScheme = (colorScheme) => {
        document.documentElement.classList.toggle("color-scheme-light", colorScheme === "light")
        document.documentElement.classList.toggle("color-scheme-dark", colorScheme === "dark")
        colorSchemeButtons.forEach((button) => {
          button.setAttribute("aria-checked", button.dataset.colorScheme === colorScheme)
        })
      }

      applyColorScheme(selectedColorScheme)

      colorSchemeMenu.addEventListener("click", (event) => {
        const colorSchemeButton = event.target.closest("[data-color-scheme]")
        if (!colorSchemeButton) return

        const colorScheme = colorSchemeButton.dataset.colorScheme
        applyColorScheme(colorScheme)

        try {
          localStorage.setItem(colorSchemeStorageKey, colorScheme)
        } catch (_) {
          // Local storage may be unavailable in privacy-restricted browsers.
        }

        if (colorSchemeMenu.hidePopover) {
          colorSchemeMenu.hidePopover()
        }
      })
    }

    const themingExampleForm = document.getElementById("theming-example-form")
    if (themingExampleForm) {
      const accentColorRange = document.getElementById('accent-color-range')

      const updateAccentHue = (hue) => {
        themingExampleForm.style.setProperty("--accent-hue", hue);
        accentColorRange.style.setProperty("--accent-hue", hue);
        accentColorRange.parentElement.dataset.value = hue
      }

      accentColorRange.addEventListener('input', (evt) => {
        updateAccentHue(evt.target.value)
      })

      const accentSaturationRange = document.getElementById('accent-saturation-range')

      const updateAccentSaturation = (saturation) => {
        themingExampleForm.style.setProperty("--accent-saturation", `${saturation}%`);
        accentColorRange.style.setProperty("--accent-saturation", `${saturation}%`);
        accentSaturationRange.parentElement.dataset.value = `${saturation}%`
      }

      accentSaturationRange.addEventListener('input', (evt) => {
        updateAccentSaturation(evt.target.value)
      })

      accentColorRange.parentElement.dataset.value = accentColorRange.value
      accentSaturationRange.parentElement.dataset.value = `${accentSaturationRange.value}%`

      const neutralColorRange = document.getElementById('neutral-color-range')

      const updateNeutralHue = (hue) => {
        themingExampleForm.style.setProperty("--neutral-hue", hue);
        neutralColorRange.style.setProperty("--accent-hue", hue);
        neutralColorRange.parentElement.dataset.value = hue
      }

      neutralColorRange.addEventListener('input', (evt) => {
        updateNeutralHue(evt.target.value)
      })

      const neutralSaturationRange = document.getElementById('neutral-saturation-range')

      const updateneutralSaturation = (saturation) => {
        themingExampleForm.style.setProperty("--neutral-saturation", `${saturation}%`);
        neutralColorRange.style.setProperty("--accent-saturation", `${saturation}%`);
        neutralSaturationRange.parentElement.dataset.value = `${saturation}%`
      }

      neutralSaturationRange.addEventListener('input', (evt) => {
        updateneutralSaturation(evt.target.value)
      })

      neutralColorRange.parentElement.dataset.value = neutralColorRange.value
      neutralSaturationRange.parentElement.dataset.value = `${neutralSaturationRange.value}%`
    }

    document.querySelectorAll("[data-highlight]").forEach((element) => {
      hljs.highlightElement(element)
    })

    const indeterminateCheckbox = document.getElementById('indeterminate-checkbox')
    if (indeterminateCheckbox) {
      indeterminateCheckbox.indeterminate = true;
    }

    const releaseArticle = document.getElementById("module-picker-article")
    if (releaseArticle) {
      const snippet = releaseArticle.querySelector(".source code")
      const moduleStorageKey = "microbe-release-modules"
      const moduleCheckboxes = releaseArticle.querySelectorAll('input[name="modules"]')
      const selectedModules = () => Array.from(new Set(Array.from(moduleCheckboxes)
        .filter((checkbox) => checkbox.checked)
        .map((checkbox) => checkbox.value)))
      const selectedChoices = () => Array.from(moduleCheckboxes)
        .filter((checkbox) => checkbox.checked)
        .map((checkbox) => checkbox.dataset.choice || checkbox.value)

      const rememberSelectedModules = () => {
        try {
          sessionStorage.setItem(moduleStorageKey, JSON.stringify(selectedChoices()))
        } catch (_) {
          // Session storage may be unavailable in privacy-restricted browsers.
        }
      }

      const restoreSelectedModules = () => {
        try {
          const storedModules = JSON.parse(sessionStorage.getItem(moduleStorageKey) || "[]")
          if (!Array.isArray(storedModules)) return

          moduleCheckboxes.forEach((checkbox) => {
            checkbox.checked = storedModules.includes(checkbox.value) ||
              storedModules.includes(checkbox.dataset.choice)
          })
        } catch (_) {
          // Ignore unavailable or malformed session storage.
        }
      }

      const updateSnippet = () => {
        snippet.textContent = releaseTemplate(releaseArticle.dataset.version, selectedModules())
        snippet.removeAttribute("data-highlighted")
        hljs.highlightElement(snippet)
      }

      releaseArticle.addEventListener("change", (event) => {
        if (event.target.matches('input[name="modules"]')) {
          rememberSelectedModules()
          updateSnippet()
        }
      })

      restoreSelectedModules()
      updateSnippet()
    }
  })

  window.addEventListener('click', (e) => {
    if (e.target.matches('button.show-source')) {
      e.target.ariaPressed = !e.target.matches('[aria-pressed=true]')
      const card = e.target.closest(".card")
      let source = card.querySelector('.source')

      if (source) {
        source.remove()
      } else {
        source = document.createElement('section')
        source.classList.add('source')

        const pre = document.createElement('pre')

        const code = document.createElement('code')
        code.classList.add('language-html')

        source.appendChild(pre)
        pre.appendChild(code)

        const snippet = document.createTextNode(card.querySelector('section').innerHTML)
        code.appendChild(snippet)

        card.appendChild(source)
        hljs.highlightElement(code);
      }
    } else if (e.target.matches('button.jsfiddle')) {
      const example = e.target.closest(".example")
      const card = example.querySelector(".card")
      const sourceSection = card.querySelector('section')
      let source = sourceSection.innerHTML

      source = wrapWithClasses(source, ['microbe', ...sourceSection.classList])

      const form = example.querySelector("form.jsfiddle")

      const modules = (form.dataset.modules || '').split(/\s+/).filter(Boolean)
      const htmlInput = form.querySelector("input[name=html]")
      htmlInput.value = htmlTemplate(form.dataset.version, modules, source)

      const cssInput = form.querySelector("input[name=css]")
      cssInput.value = cssTemplate

      form.submit()
    } else if (e.target.matches('button.copy-source')) {
      const card = e.target.closest(".card")
      let source = card.querySelector('.source code')
      navigator.clipboard.writeText(source.textContent)
    } else if (e.target.matches('button.color-sample')) {
      const { hue, saturation, lightness } = e.target.dataset
      navigator.clipboard.writeText(`hsl(${hue} ${saturation} var(--lightness-${lightness}))`)
    }
  })

  window.addEventListener("keydown", (event) => {
    if (event.key === "k" && (event.metaKey || event.ctrlKey)) {
      event.preventDefault();

      const dialog = document.getElementById("search-dialog");
      if (!dialog) return;

      if (!dialog.open) {
        dialog.showModal();
      }
    }
  });
}
