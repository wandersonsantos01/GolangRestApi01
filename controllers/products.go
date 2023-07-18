package controllers

import (
	"golang_rest_api/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var templ = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	// products := []models.Product{
	// 	{Name: "Camiseta", Description: "Camiseta verde", Price: 98.99, Quantity: 15},
	// 	{"Tenis", "Confortável", 199.98, 5},
	// 	{"Headset", "JBL", 200, 2},
	// 	{"T-Shirt", "Red", 99, 155},
	// }

	products := models.GetAllProducts()

	templ.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Convert error - price")
		}

		convertedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Convert error - quantity")
		}

		models.CreateNewProduct(name, description, convertedPrice, convertedQuantity)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	models.DeleteProduct(id)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	produto := models.GetProductById(id)

	templ.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Convert error - price")
		}

		convertedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Convert error - quantity")
		}

		models.UpdateProduct(id, name, description, convertedPrice, convertedQuantity)
	}

	http.Redirect(w, r, "/", 301)
}
