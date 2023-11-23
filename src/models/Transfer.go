package models

import (
	// "github.com/LeonardoMuller13/digital-bank-api/src/models"
	"gorm.io/gorm"
)

type Transfer struct {
	gorm.Model
	Account_origin_ID      int
	Account_destination_ID int
	Amount                 int
}

var Transfers []Transfer
