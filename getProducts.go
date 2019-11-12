package main

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []Product

	result, err := db.Query("SELECT * from product")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var product Product

		err := result.Scan(&product.ProductID, &product.Title, &product.Sku, &product.Description, &product.Price, &product.Created, &product.LastUpdated)
		if err != nil {
			panic(err.Error())
		}

		products = append(products, product)
	}

	json.NewEncoder(w).Encode(products)
}
