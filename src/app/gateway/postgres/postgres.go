package postgres

import (
	"context"
	"log"

	"github.com/LeonardoMuller13/digital-bank-api/src/app/config"
	"github.com/LeonardoMuller13/digital-bank-api/src/app/domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Client struct {
	DB *gorm.DB
}

func (c *Client) Close() {
	dbInstance, _ := c.DB.DB()
	_ = dbInstance.Close()
}

// New connects to the Postgres database and performs migrations.
func New(ctx context.Context, config config.Postgres) (*Client, error) {
	const (
		operation = "Postgres.New"
	)

	connString := "host=" + config.Host + " user=" + config.User + " password=" + config.Password + " dbname=" + config.DatabaseName + " port=" + config.Port + " sslmode=disable"
	DB, err := gorm.Open(postgres.Open(connString))

	if err != nil {
		log.Fatalf("%s -> %w", operation, err)
		return nil, err
	}

	if !DB.Migrator().HasTable(&entity.Account{}) {
		DB.Migrator().CreateTable(&entity.Account{})
	}
	if !DB.Migrator().HasTable(&entity.Transfer{}) {
		DB.Migrator().CreateTable(&entity.Transfer{})
	}

	return &Client{DB}, nil
}
