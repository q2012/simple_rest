package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"simple_rest/api"
	"simple_rest/globals"
	"simple_rest/models"
)


var arr[] models.Main

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func initApi() *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/all", api.AddHeaders(api.GetAll, "application/json")).Methods("GET")
	r.HandleFunc("/all/{id}", api.AddHeaders(api.GetOne, "application/json")).Methods("GET")
	r.HandleFunc("/all", api.AddHeaders(api.CreateOne, "application/json")).Methods("POST")
	//r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	//r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	return r
}

func main() {

	globals.InitGlobals()

	r := initApi()
	log.Fatal(http.ListenAndServe(":8000", r))
}
