package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "<user>:<password>@tcp(127.0.0.1:3306)/<database>")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/api/products", getProducts).Methods("GET")
	router.HandleFunc("/api/products/{productId}", getProduct).Methods("GET")
	router.HandleFunc("/api/products", createProduct).Methods("POST")
	router.HandleFunc("/api/products/{productId}", updateProduct).Methods("PUT")
	router.HandleFunc("/api/products/{productId}", deleteProduct).Methods("DELETE")
	http.ListenAndServe(":8000", router)
}
