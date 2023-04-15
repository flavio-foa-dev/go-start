package controllers

import (
	"app/apiRestGo/db"
	"app/apiRestGo/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page")

}

func GetPersonAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var persons []models.Person
	db.DB.Find(&persons)
	json.NewEncoder(w).Encode(persons)
}

func GetPersonById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	var person models.Person
	db.DB.First(&person, id)

	json.NewEncoder(w).Encode(person)

}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	json.NewDecoder(r.Body).Decode(&person)
	db.DB.Create(&person)

	json.NewEncoder(w).Encode(person)
}

func DestroyPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var person models.Person

	db.DB.Delete(&person, id)
	json.NewEncoder(w).Encode(person)

}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var person models.Person

	db.DB.First(&person, id)
	json.NewDecoder(r.Body).Decode(&person)

	db.DB.Save(&person)

	json.NewEncoder(w).Encode(&person)
}
