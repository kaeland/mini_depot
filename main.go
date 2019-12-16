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
	w.Header().Set("Content-Type", "application/json")

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

	log.Fatal(http.ListenAndServe(":8000", r))
}
