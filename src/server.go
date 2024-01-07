package main

import (
	"fmt"
	"log"

	"github.com/LeonardoMuller13/digital-bank-api/src/config"
	"github.com/LeonardoMuller13/digital-bank-api/src/database"
	"github.com/LeonardoMuller13/digital-bank-api/src/routes"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configurations: %v", err)
	}

	err = database.ConnectDB(cfg.Postgres)
	if err != nil {
		log.Fatalf("Failed to start postgres: %v", err)
	}

	fmt.Println("STARTING SERVER")
	routes.HandleRequest()
}
