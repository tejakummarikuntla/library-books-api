package main

import (
	"book-store/controllers"
	"book-store/driver"
	"book-store/models"
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
)

var books []models.Book
var db *sql.DB

func logFatal (err error){
	if err != nil {
		log.Fatal(err)
	}
}

func init(){
	gotenv.Load()
}

func main() {
// connect DB from .env of Host environmental variables
	db = driver.ConnectDB()

	controller := controllers.Controller{}

	router := mux.NewRouter()
	router.HandleFunc("/books",controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}",controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books",controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books",controller.UpdateBooks(db)).Methods("PUT")
	router.HandleFunc("/books/{id}",controller.RemoveBook(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000",router))
}
