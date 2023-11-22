package routes

import (
	"log"
	"net/http"

	"github.com/LeonardoMuller13/digital-bank-api/controllers"
	"github.com/LeonardoMuller13/digital-bank-api/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	s := r.PathPrefix("/accounts").Subrouter()
	s.Use(middleware.ProtectedHandler)
	s.HandleFunc("/", controllers.GetAccounts).Methods("GET")
	s.HandleFunc("/", controllers.NewAccount).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
