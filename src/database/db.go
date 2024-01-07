package database

import (
	"github.com/LeonardoMuller13/digital-bank-api/src/config"
	"github.com/LeonardoMuller13/digital-bank-api/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB(config config.Postgres) error {
	connection := "host=" + config.Host + " user=" + config.User + " password=" + config.Password + " dbname=" + config.DatabaseName + " port=" + config.Port + " sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connection))

	if err != nil {
		return err
	}

	if !DB.Migrator().HasTable(&models.Account{}) {
		DB.Migrator().CreateTable(&models.Account{})
	}
	if !DB.Migrator().HasTable(&models.Transfer{}) {
		DB.Migrator().CreateTable(&models.Transfer{})
	}
	return nil
}
