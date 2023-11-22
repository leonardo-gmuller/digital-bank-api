package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name    string
	Cpf     string
	Secret  string
	Balance int
}

var Accounts []Account
