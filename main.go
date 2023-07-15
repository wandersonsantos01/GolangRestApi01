package main

import (
	"golang_rest_api/models"
	"net/http"
	"text/template"
)

var templ = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	// products := []models.Product{
	// 	{Name: "Camiseta", Description: "Camiseta verde", Price: 98.99, Quantity: 15},
	// 	{"Tenis", "Confortável", 199.98, 5},
	// 	{"Headset", "JBL", 200, 2},
	// 	{"T-Shirt", "Red", 99, 155},
	// }

	products := models.GetAllProducts()

	templ.ExecuteTemplate(w, "Index", products)
}
