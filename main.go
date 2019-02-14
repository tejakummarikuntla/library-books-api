package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"os"
)

type Book struct {
	ID      int     `json:id`
	Title   string  `json:title`
	Author  string  `json:author`
	Year    string  `json:year`
}

var books []Book
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

	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)

	db, err = sql.Open("postgres", pgUrl)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	log.Println(pgUrl)

	router := mux.NewRouter()
	router.HandleFunc("/books",getBooks).Methods("GET")
	router.HandleFunc("/books/{id}",getBook).Methods("GET")
	router.HandleFunc("/books",addBook).Methods("POST")
	router.HandleFunc("/books",updateBooks).Methods("PUT")
	router.HandleFunc("/books/{id}",removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000",router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	var book Book
	books = []Book{}

	rows, err := db.Query("SELECT * FROM books")
	logFatal(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)
		books = append(books, book)
	}
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	params := mux.Vars(r)
	/*books = []Book{}
	prams := mux.Vars(r)
	id, err := strconv.Atoi(prams["id"])
	logFatal(err)
	rows, err := db.Query("SELECT * FROM books")
	logFatal(err)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)
		books = append(books, book)
		if book.ID == id{
			json.NewEncoder(w).Encode(book)
			return
		}
	}*/
	rows := db.QueryRow("SELECT  * FROM books WHERE id=$1", params["id"])
	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	logFatal(err)
	json.NewEncoder(w).Encode(book)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	var bookID int

	json.NewDecoder(r.Body).Decode(&book)
	err := db.QueryRow("INSERT INTO books (title, author, year) values ($1, $2, $3) RETURNING id;",
						book.Title,
						book.Author,
						book.Year,
	).Scan(&bookID)
	logFatal(err)

	json.NewEncoder(w).Encode(bookID)
}

func updateBooks(w http.ResponseWriter, r *http.Request) {

}

func removeBook(w http.ResponseWriter, r *http.Request){


}
