if (window.self === window.top) {
  window.addEventListener('load', () => {
    const themingExampleForm = document.getElementById("theming-example-form")
    if (themingExampleForm) {
      const colorRange = document.getElementById('color-range')

      const updateAccentHue = (hue) => {
        themingExampleForm.style.setProperty("--accent-hue", hue);
        colorRange.style.setProperty("--accent-hue", hue);
      }

      colorRange.addEventListener('input', (evt) => {
        updateAccentHue(evt.target.value)
      })

      const saturationRange = document.getElementById('saturation-range')

      const updateAccentSaturation = (saturation) => {
        themingExampleForm.style.setProperty("--accent-saturation", `${saturation}%`);
        colorRange.style.setProperty("--accent-saturation", `${saturation}%`);
      }

      saturationRange.addEventListener('input', (evt) => {
        updateAccentSaturation(evt.target.value)
      })
    }

    const indeterminateCheckbox = document.getElementById('indeterminate-checkbox')
    if (indeterminateCheckbox) {
      indeterminateCheckbox.indeterminate = true;
    }

    if (typeof hljs !== "undefined") {
      const codeBlocks = document.querySelectorAll("[data-highlight=yes]")
      console.log(codeBlocks)
      for (const codeBlock of codeBlocks) {
        hljs.highlightElement(codeBlock);
      }
    }
  })
}
