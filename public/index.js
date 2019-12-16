const baseUrl = "http://localhost:8000/api/products"

document.addEventListener("DOMContentLoaded", init)

function init() {
  fetchProducts()
}

function fetchProducts() {
  fetch(baseUrl)
    .then(res => res.json())
    .then(products => renderProducts(products))
}

function renderProducts(products) {
  let productsContainer = document.querySelector(".products-container")

  products.map((product) => {
    let { Code, Price } = product

    let div = document.createElement("div")
    div.className = "product-item"

    let productCode = document.createElement("p")
    productCode.className = "product-property"

    let productPrice = document.createElement("p")
    productPrice.className = "product-property" 
    
    productCode.innerText = `Code: ${Code}`
    productPrice.innerText = `Price: $${Price * 0.01}`

    div.appendChild(productCode)
    div.appendChild(productPrice)
    productsContainer.appendChild(div)
  })
}