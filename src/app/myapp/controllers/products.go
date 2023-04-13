package controllers

import (
	"app/myapp/models"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var views = template.Must(template.ParseGlob("./templates/*.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	products := models.GetProductsAll()

	views.ExecuteTemplate(w, "Index", products)
}

func CreateFormProduct(w http.ResponseWriter, r *http.Request) {
	views.ExecuteTemplate(w, "Cadastrar", nil)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		quantity := r.FormValue("quantidade")

		parsePrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Error na conversao", err)
		}

		parseQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			fmt.Println("Error na conversao", err)
		}
		models.InsertProduct(name, description, parsePrice, parseQuantity)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")

	models.DeleteProduct(idProduct)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func UpdateForm(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	product := models.EditProduct(idProduct)

	views.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		quantity := r.FormValue("quantidade")

		parseId, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error na conversao", err)
		}

		parsePrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error na conversao", err)
		}

		parseQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error na conversao", err)
		}
		models.UpdateProduct(parseId, name, description, parsePrice, parseQuantity)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
