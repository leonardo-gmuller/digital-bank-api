package tests

import (
	"testing"

	"github.com/LeonardoMuller13/digital-bank-api/src/config"
	"github.com/LeonardoMuller13/digital-bank-api/src/database"
	"github.com/LeonardoMuller13/digital-bank-api/src/models"
	"gorm.io/gorm"
)

func setupSuite(tb testing.TB) func(tb testing.TB) {
	cfg, err := config.Load()
	if err != nil {
		tb.Fatalf("failed to load configurations: %v", err)
	}
	database.ConnectDB(cfg.Postgres)
	// Return a function to teardown the test
	return func(tb testing.TB) {
		database.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Account{})
		database.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Transfer{})
	}
}
