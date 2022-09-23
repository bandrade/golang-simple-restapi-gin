package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name string `json:"name"`
	RG   string `json:"rg"`
	CPF  string `json:"cpf"`
}

var Students []Student
