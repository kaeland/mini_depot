package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Product model
type Product struct {
	gorm.Model
	Code     string
	Price    uint
	ImageURL string
}

var db *gorm.DB
var err error

func getProducts(w http.ResponseWriter, r *http.Request) {
	var products []Product

	w.Header().Set("Content-Type", "application/json")

	db.Find(&products)
	json.NewEncoder(w).Encode(&products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	var params = mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	db.First(&product, params["id"])
	json.NewEncoder(w).Encode(&product)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var product Product

	w.Header().Set("Content-Type", "application/json")

	json.NewDecoder(r.Body).Decode(&product)
	db.Create(&product)
	json.NewEncoder(w).Encode(&product)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	var params = mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	db.First(&product, params["id"])
	db.Update(&product)

	json.NewDecoder(r.Body).Decode(&product)
	json.NewEncoder(w).Encode(&product)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	var params = mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")
	db.First(&product, params["id"])
	db.Delete(&product)

	var products []Product
	db.Find(&products)
	json.NewEncoder(w).Encode(&products)
}

func main() {
	r := mux.NewRouter()

	db, err = gorm.Open("sqlite3", "mini_depot.db")
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(&Product{})

	// product1 := Product{Code: "E234", Price: 1800, ImageURL: "images.com/floorboards"}
	// product2 := Product{Code: "E123", Price: 1500, ImageURL: "images.com/shovels"}
	// product3 := Product{Code: "E687", Price: 450, ImageURL: "images.com/toolset"}

	r.HandleFunc("/api/products", getProducts).Methods("GET")
	r.HandleFunc("/api/products/{id}", getProduct).Methods("GET")
	r.HandleFunc("/api/products", createProduct).Methods("POST")
	r.HandleFunc("/api/products/{id}", updateProduct).Methods("PUT")
	r.HandleFunc("/api/products/{id}", deleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
