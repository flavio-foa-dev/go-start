package routes

import (
	"app/apiRestGo/controllers"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.Home)
	router.HandleFunc("/api/person", controllers.GetPersonAll).Methods("Get")
	router.HandleFunc("/api/person/{id}", controllers.GetPersonById).Methods("Get")

	router.HandleFunc("/api/person", controllers.CreatePerson).Methods("Post")

	router.HandleFunc("/api/person/{id}", controllers.UpdatePerson).Methods("Put")

	router.HandleFunc("/api/person/{id}", controllers.DestroyPerson).Methods("Delete")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router)))
}
