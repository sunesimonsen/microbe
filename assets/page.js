if (window.self === window.top) {
  window.addEventListener('load', () => {
    const colorRange = document.getElementById('color-range')
    if (colorRange) {
      const saturationRange = document.getElementById('saturation-range')
      const themingExampleForm = document.getElementById("theming-example-form")

      const updateAccentHue = (hue) => {
        themingExampleForm.style.setProperty("--accent-hue", hue);
        colorRange.style.setProperty("--accent-hue", hue);
      }

      colorRange.addEventListener('input', (evt) => {
        updateAccentHue(evt.target.value)
      })

      const updateAccentSaturation = (saturation) => {
        themingExampleForm.style.setProperty("--accent-saturation", `${saturation}%`);
        colorRange.style.setProperty("--accent-saturation", `${saturation}%`);
      }

      saturationRange.addEventListener('input', (evt) => {
        updateAccentSaturation(evt.target.value)
      })
    }
  })
}
