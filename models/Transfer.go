package models

import (
	"gorm.io/gorm"
)

type Transfer struct {
	gorm.Model
	Account_Origin_ID      int
	Account_Destination_ID int
	Amount                 int
}

var Transfers []Transfer
