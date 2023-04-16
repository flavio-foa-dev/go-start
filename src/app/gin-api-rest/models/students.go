package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Nome  string `json:"nome"`
	CPF   string `json:"cpf"`
	RG    string `json:"rg"`
	Email string `json:"email"`
}
