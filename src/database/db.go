package database

import (
	"log"
	"os"

	"github.com/LeonardoMuller13/digital-bank-api/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	stringDeConexao := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable"
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
