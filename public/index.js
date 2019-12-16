const baseUrl = "http://localhost:8000/api/products"

document.addEventListener("DOMContentLoaded", init)

function init() {
  fetchProducts()
}

function fetchProducts() {
  fetch(baseUrl)
    .then(res => res.json())
    .then(console.log)
}