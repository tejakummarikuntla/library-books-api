package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Book struct {
	ID      int     `json:id`
	Title   string  `json:title`
	Author  string  `json:author`
	Year    string  `json:year`
}

var Books []Book

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/books",getBooks).Methods("GET")
	router.HandleFunc("/books/{id}",getBook).Methods("GET")
	router.HandleFunc("/books",addBook).Methods("POST")
	router.HandleFunc("/books",updateBooks).Methods("PUT")
	router.HandleFunc("/books/{id}",removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000",router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Gets all the books")
}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Get one the books")
}

func addBook(w http.ResponseWriter, r *http.Request)  {
	log.Println("Adds Books")
}

func updateBooks(w http.ResponseWriter, r *http.Request)  {
	log.Println("Updates a book")
}

func removeBook(w http.ResponseWriter, r *http.Request){
	log.Println("Delets a Book")
}
