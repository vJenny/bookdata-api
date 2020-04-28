package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matt-FFFFFF/bookdata-api/datastore"
)

var (
	books datastore.BookStore
)

func init() {
	books = &datastore.Books{}
	books.Initialize()
}

func main() {
	r := mux.NewRouter()
	log.Println("bookdata api")
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "api v1")
	})
	api.HandleFunc("/books", getAllBooks).Methods(http.MethodGet)
	api.HandleFunc("/books/authors/{author}", getAllBooksByAuthor).Methods(http.MethodGet)
	api.HandleFunc("/books/title/{title}", getAllBooksByTitle).Methods(http.MethodGet)
	api.HandleFunc("/book/isbn/{isbn}", getBookByISBN).Methods(http.MethodGet)
	api.HandleFunc("/book", addBookByISBN).Methods(http.MethodPost)
	api.HandleFunc("/book/isbn/{isbn}", removeBookByISBN).Methods(http.MethodDelete)
	log.Fatalln(http.ListenAndServe(":8080", r))
}
