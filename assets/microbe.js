
window.addEventListener('click', (e) => {
  console.log(e.target)


  if (e.target.matches('.menu-toggle')) {
    document.body.classList.toggle('show-menu')
  } else if (e.target.matches('a') && document.body.classList.contains('show-menu')) {
    document.body.classList.remove('show-menu')
  }
})
