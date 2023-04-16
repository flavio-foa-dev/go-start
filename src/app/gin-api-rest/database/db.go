package database

import (
	"api-gin/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connection() {

	dsn := "host=localhost user=root password=docker dbname=alunos port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic("Erro na conexao com db", err.Error())
	}

	DB.AutoMigrate(&models.Student{})

}
