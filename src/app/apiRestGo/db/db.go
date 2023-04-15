package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connection() {
	dsn := "host=localhost user=root password=docker dbname=postgres port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))
	log.Printf("Conectando com o servidor")

	if err != nil {
		log.Panic("Error connecting with  database:", err.Error())
	}
}
