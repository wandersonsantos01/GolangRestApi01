package models

import (
	"golang_rest_api/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAllProducts() []Product {
	db := db.DbConnect()
	rows, err := db.Query("SELECT * FROM products ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}
	for rows.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = rows.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func CreateNewProduct(name, description string, price float64, quantity int) {
	db := db.DbConnect()

	insertQry, err := db.Prepare("INSERT INTO products (name, description, price, quantity) VALUES ($1, $2, $3, $4);")
	if err != nil {
		panic(err.Error())
	}

	insertQry.Exec(name, description, price, quantity)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.DbConnect()

	deleteQry, err := db.Prepare("DELETE FROM products WHERE id = $1;")
	if err != nil {
		panic(err.Error())
	}

	deleteQry.Exec(id)
	defer db.Close()
}

func GetProductById(id string) Product {
	db := db.DbConnect()

	getProduct, err := db.Query("SELECT * FROM products WHERE id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	product := Product{}

	for getProduct.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = getProduct.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity
	}

	defer db.Close()

	return product
}

func UpdateProduct(id, name, description string, price float64, quantity int) {
	db := db.DbConnect()

	editQry, err := db.Prepare("UPDATE products SET name = $1, description = $2, price = $3, quantity = $4 WHERE id = $5;")
	if err != nil {
		panic(err.Error())
	}

	editQry.Exec(name, description, price, quantity, id)
	defer db.Close()
}
