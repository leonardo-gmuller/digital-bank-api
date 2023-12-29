package database

import (
	"log"
	"os"

	"github.com/LeonardoMuller13/digital-bank-api/src/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	connection := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connection))

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
