package routes

import (
	"app/myapp/controllers"
	"net/http"
)

func Router() {
	http.HandleFunc("/", controllers.IndexHandler)
	http.HandleFunc("/cadastrar", controllers.CreateFormProduct)
	http.HandleFunc("/insert", controllers.CreateProduct)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.UpdateForm)
	http.HandleFunc("/update", controllers.Update)
}
