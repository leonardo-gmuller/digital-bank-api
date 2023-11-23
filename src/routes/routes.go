package routes

import (
	"log"
	"net/http"

	"github.com/LeonardoMuller13/digital-bank-api/src/controllers"
	"github.com/LeonardoMuller13/digital-bank-api/src/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	s := r.PathPrefix("/accounts").Subrouter()
	s.HandleFunc("", controllers.GetAccounts).Methods("GET")
	s.HandleFunc("", controllers.NewAccount).Methods("POST")
	s.HandleFunc("/{account_id}/balance", controllers.GetAccountBalanceByID).Methods("GET")
	s2 := r.PathPrefix("/transfers").Subrouter()
	s2.HandleFunc("", controllers.GetTransfers).Methods("GET")
	s2.HandleFunc("", controllers.NewTransfer).Methods("POST")
	s2.Use(middleware.ProtectedHandler)
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
