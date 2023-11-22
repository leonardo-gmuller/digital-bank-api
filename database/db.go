package database

import (
	"log"

	"github.com/LeonardoMuller13/digital-bank-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))

	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados.")
	}

	if !DB.Migrator().HasTable(&models.Account{}) {
		DB.Migrator().CreateTable(&models.Account{})
	}
	if !DB.Migrator().HasTable(&models.Transfer{}) {
		DB.Migrator().CreateTable(&models.Transfer{})
	}
}
