package datastore

import "github.com/matt-FFFFFF/bookdata-api/loader"

// BookStore is the interface that the http methods use to call the backend datastore
// Using an interface means we could replace the datastore with something else,
// as long as that something else provides these method signatures...
type BookStore interface {
	Initialize()
	GetAllBooks(limit, skip int) *[]*loader.BookData
	GetAllBooksByAuthor(author string) *[]*loader.BookData
	GetAllBooksByTitle(title string) *[]*loader.BookData
	GetBookByISBN(isbn string) *loader.BookData
	RemoveBookByISBN(isbn string) bool
	AddBook(book loader.BookData)
}
