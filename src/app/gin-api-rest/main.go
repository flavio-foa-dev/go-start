package main

import (
	"api-gin/database"
	"api-gin/routes"
)

func main() {
	database.Connection()
	routes.HandleRequest()
}
