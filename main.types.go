package main

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

type Product struct {
	ProductID   int    `json:"productId"`
	Title       string `json:"title"`
	Sku         string `json:"sku"`
	Barcodes    []Barcode
	Description sql.NullString `json:"description"`
	Attributes  []ProductAttribute
	Price       float32        `json:"price"`
	Created     mysql.NullTime `json:"created"`
	LastUpdated mysql.NullTime `json:"lastUpdated"`
}

type Barcode struct {
	ProductID int    `json:"productId"`
	Barcode   string `json:"barcode"`
}

type ProductAttribute struct {
	ProductID int    `json:"productId"`
	Name      string `json:"name"`
	Value     string `json:"value"`
}
