package main

import (
	"app/myapp/routes"
	"net/http"
)

func main() {
	routes.Router()
	http.ListenAndServe(":8080", nil)

}
