package tests

import (
	"testing"

	"github.com/LeonardoMuller13/digital-bank-api/src/config"
	"github.com/LeonardoMuller13/digital-bank-api/src/database"
	"github.com/LeonardoMuller13/digital-bank-api/src/models"
	"gorm.io/gorm"
)

func setupSuite(tb testing.TB) func(tb testing.TB) {
	tb.Setenv("ENVIRONMENT", string(config.EnvTest))
	tb.Setenv("DEVELOPMENT", "false")
	tb.Setenv("JWT_SECRET", "my-secret-key")
	tb.Setenv("DB_HOST", "local")
	tb.Setenv("DB_USER", "postgres")
	tb.Setenv("DB_PASSWORD", "postgres")
	tb.Setenv("DB_NAME", "digitalbank_test")
	tb.Setenv("DB_PORT", "5432")
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
