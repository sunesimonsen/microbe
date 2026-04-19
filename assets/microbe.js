if (window.self === window.top) {
  window.addEventListener('click', (e) => {
    if (e.target.matches('.menu-toggle')) {
      document.body.classList.toggle('show-menu')
    } else if (e.target.matches('a') && document.body.classList.contains('show-menu')) {
      document.body.classList.remove('show-menu')
    }
  })
}
