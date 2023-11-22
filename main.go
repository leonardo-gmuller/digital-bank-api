package main

import (
	"fmt"

	"github.com/LeonardoMuller13/digital-bank-api/database"
	"github.com/LeonardoMuller13/digital-bank-api/routes"
)

func main() {
	database.ConnectDB()
	fmt.Println("STARTING SERVER")
	routes.HandleRequest()
}
