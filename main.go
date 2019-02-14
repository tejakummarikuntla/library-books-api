package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

type Book struct {
	ID      int     `json:id`
	Title   string  `json:title`
	Author  string  `json:author`
	Year    string  `json:year`
}

var books []Book

func main() {
	router := mux.NewRouter()

	books = append(books,
				Book{ID: 1, Title: "title-1" , Author:"author-1", Year:"year-1"},
				Book{ID: 2, Title: "title-2" , Author:"author-2", Year:"year-2"},
				Book{ID: 3, Title: "title-3" , Author:"author-3", Year:"year-3"},
				Book{ID: 4, Title: "title-4" , Author:"author-4", Year:"year-4"},
				Book{ID: 5, Title: "title-5" , Author:"author-5", Year:"year-5"},
	)

	router.HandleFunc("/books",getBooks).Methods("GET")
	router.HandleFunc("/books/{id}",getBook).Methods("GET")
	router.HandleFunc("/books",addBook).Methods("POST")
	router.HandleFunc("/books",updateBooks).Methods("PUT")
	router.HandleFunc("/books/{id}",removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000",router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Gets all the books")

	json.NewEncoder(w).Encode(books)

}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Get one the books")
	params := mux.Vars(r)
	log.Println(reflect.TypeOf(params))

	i, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, book := range books {
		if book.ID == i{
			json.NewEncoder(w).Encode(book)
		}
	}
}

func addBook(w http.ResponseWriter, r *http.Request)  {
	log.Println("Adds Books")

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)

	json.NewEncoder(w).Encode(books)
}

func updateBooks(w http.ResponseWriter, r *http.Request)  {
	log.Println("Updates a book")
}

func removeBook(w http.ResponseWriter, r *http.Request){
	log.Println("Delets a Book")
}
