package models

import (
	"time"

	"gorm.io/gorm"
)

type Transfer struct {
	ID                   uint           `gorm:"primaryKey";json:"id"`
	AccountOriginId      int            `json:"account_origin_id"`
	AccountDestinationId int            `json:"account_destination_id"`
	Amount               int            `json:"amount"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
	DeletedAt            gorm.DeletedAt `gorm:"index";json:"deleted_at"`
}

var Transfers []Transfer
