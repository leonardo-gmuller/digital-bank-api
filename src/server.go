package main

import (
	"fmt"

	"github.com/LeonardoMuller13/digital-bank-api/src/database"
	"github.com/LeonardoMuller13/digital-bank-api/src/routes"
)

func main() {
	database.ConnectDB()
	fmt.Println("STARTING SERVER")
	routes.HandleRequest()
}
