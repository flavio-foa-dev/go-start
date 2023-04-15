package main

import (
	"app/apiRestGo/db"
	"app/apiRestGo/models"
	"app/apiRestGo/routes"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the")
}

func exemplo() {
	models.Persons = []models.Person{
		{Id: 1, Name: "John", History: "1"},
		{Id: 2, Name: "John wick", History: "2"},
		{Id: 3, Name: "John Macarnner", History: "3"},
	}

}

func main() {
	exemplo()
	db.Connection()
	fmt.Println("Iniciando Sistema")
	routes.HandleRequest()
}
