const htmlTemplate = (source) => `\
<link href="https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@HEAD/assets/microbe.css" rel="stylesheet" type="text/css">
<link href="https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@HEAD/assets/microbe-accordion.css" rel="stylesheet" type="text/css">
<link href="https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@HEAD/assets/microbe-card.css" rel="stylesheet" type="text/css">
<link href="https://cdn.jsdelivr.net/gh/sunesimonsen/microbe@HEAD/assets/microbe-navlist.css" rel="stylesheet" type="text/css">
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
      const colorRange = document.getElementById('color-range')

      const updateAccentHue = (hue) => {
        themingExampleForm.style.setProperty("--accent-hue", hue);
        colorRange.style.setProperty("--accent-hue", hue);
        colorRange.parentElement.dataset.value = hue
      }

      colorRange.addEventListener('input', (evt) => {
        updateAccentHue(evt.target.value)
      })

      const saturationRange = document.getElementById('saturation-range')

      const updateAccentSaturation = (saturation) => {
        themingExampleForm.style.setProperty("--accent-saturation", `${saturation}%`);
        colorRange.style.setProperty("--accent-saturation", `${saturation}%`);
        saturationRange.parentElement.dataset.value = `${saturation}%`
      }

      saturationRange.addEventListener('input', (evt) => {
        updateAccentSaturation(evt.target.value)
      })

      colorRange.parentElement.dataset.value = colorRange.value
      saturationRange.parentElement.dataset.value = `${saturationRange.value}%`
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
    }
  })
}
