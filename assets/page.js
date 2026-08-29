const htmlTemplate = (source) => `\
<link href="https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@HEAD/assets/microbe.css" rel="stylesheet" type="text/css">
<link href="https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@HEAD/assets/microbe-accordion.css" rel="stylesheet" type="text/css">
<link href="https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@HEAD/assets/microbe-avatar.css" rel="stylesheet" type="text/css">
<link href="https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@HEAD/assets/microbe-breadcrumb.css" rel="stylesheet" type="text/css">
<link href="https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@HEAD/assets/microbe-card.css" rel="stylesheet" type="text/css">
<link href="https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@HEAD/assets/microbe-dialog.css" rel="stylesheet" type="text/css">
<link href="https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@HEAD/assets/microbe-navlist.css" rel="stylesheet" type="text/css">
<link href="https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@HEAD/assets/microbe-pagination.css" rel="stylesheet" type="text/css">
<link href="https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@HEAD/assets/microbe-skeleton.css" rel="stylesheet" type="text/css">
<link href="https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@HEAD/assets/microbe-tabs.css" rel="stylesheet" type="text/css">

${source}
`

const indent = (source, spacing) => source.replace(/^/gm, spacing)

const wrapWithClasses = (source, classes) => `<section class="${classes.join(" ")}">
${indent(source, '  ')}
</section>`

if (window.self === window.top) {
  window.addEventListener('load', () => {
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

    const indeterminateCheckbox = document.getElementById('indeterminate-checkbox')
    if (indeterminateCheckbox) {
      indeterminateCheckbox.indeterminate = true;
    }

    for (const code of document.querySelectorAll("[data-highlight=yes]")) {
      hljs.highlightElement(code);
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

      const htmlInput = form.querySelector("input[name=html]")
      htmlInput.value = source

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
}
