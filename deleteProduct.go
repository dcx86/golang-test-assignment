package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	stmtProduct, err := db.Prepare("DELETE FROM product WHERE product_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmtProduct.Exec(params["productId"])
	if err != nil {
		panic(err.Error())
	}

	stmtProductBarcode, err := db.Prepare("DELETE FROM product_barcode WHERE product_id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmtProductBarcode.Exec(params["productId"])
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Product with ID = %s was deleted", params["productId"])
}
