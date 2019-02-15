package bookRepository

import (
	"book-store/models"
	"database/sql"
	"log"
)

type BookRepository struct {}


func logFatal (err error){
	if err != nil {
		log.Fatal(err)
	}
}

func (b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) []models.Book{
	rows, err := db.Query("SELECT * FROM books")
	logFatal(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)
		books = append(books, book)
	}
	return books
}

func (b BookRepository) GetBook(db *sql.DB, book models.Book, id int) models.Book {

	rows := db.QueryRow("SELECT  * FROM books WHERE id=$1", id)
	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	logFatal(err)

	return book
}

func (b BookRepository) AddBook(db *sql.DB,book models.Book) int{
	err := db.QueryRow("INSERT INTO books (title, author, year) values ($1, $2, $3) RETURNING id;",
		book.Title,
		book.Author,
		book.Year,
	).Scan(&book.ID)
	logFatal(err)

	return book.ID
}

func (b BookRepository) UpdateBook(db *sql.DB, book models.Book) int64{

	result, err := db.Exec("UPDATE books SET title=$1, author=$2, year=$3 WHERE id=$4 RETURNING id",
		&book.Title,
		&book.Author,
		&book.Year,
		&book.ID)

	rowsUpdated, err := result.RowsAffected()
	logFatal(err)

	return rowsUpdated
}

func (b BookRepository) RemoveBook(db *sql.DB, id int) int64 {
	result, err := db.Exec("DELETE FROM books WHERE id=$1",id)
	logFatal(err)

	rowsDeleted, err := result.RowsAffected()
	logFatal(err)
	return rowsDeleted
}