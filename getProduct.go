package main

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	result, err := db.Query("SELECT * FROM product WHERE product_id = ?", params["productId"])
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var product Product
	for result.Next() {
		err := result.Scan(&product.ProductID, &product.Title, &product.Sku, &product.Description, &product.Price, &product.Created, &product.LastUpdated)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(product)
}
