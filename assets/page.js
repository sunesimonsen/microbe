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
  })

  window.addEventListener('click', (e) => {
    if (e.target.matches('.menu-toggle')) {
      document.body.classList.toggle('show-menu')
    } else if (e.target.matches('a') && document.body.classList.contains('show-menu')) {
      document.body.classList.remove('show-menu')
    } else if (e.target.matches('.show-source')) {
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
    }
  })
}
