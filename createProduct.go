package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	stmtProduct, err := db.Prepare("INSERT INTO product(product_id, title, sku, description, price, created) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}

	stmtProductBarcode, err := db.Prepare("INSERT INTO product_barcode(product_id, barcode) VALUES(?, ?)")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	var keyVal map[string]interface{}
	if err := json.Unmarshal([]byte(body), &keyVal); err != nil {
		panic(err)
	}

	productID := 123
	barcodes := keyVal["barcodes"].([]interface{})
	// attributes := keyVal["attributes"]
	title := keyVal["title"]
	sku := keyVal["sku"]
	description := keyVal["description"]
	price := keyVal["price"]
	created := "1000-01-01 00:00:00"

	for i := 0; i < len(barcodes); i++ {
		println(barcodes[i].(string))

		barcode := barcodes[i].(string)
		_, err = stmtProductBarcode.Exec(productID, barcode)
		if err != nil {
			panic(err.Error())
		}
	}

	_, err = stmtProduct.Exec(productID, title, sku, description, price, created)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "New product was created")
}
