package datastore

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
	"time"

	"github.com/matt-FFFFFF/bookdata-api/loader"
)

// Books is the memory-backed datastore used by the API
// It contains a single field 'Store', which is (a pointer to) a slice of loader.BookData struct pointers
type Books struct {
	Store *[]*loader.BookData `json:"store"`
}

func (b *Books) strToBookData(xs []string) loader.BookData {
	return loader.BookData{
		BookID:        xs[0],
		Title:         xs[1],
		Authors:       xs[2],
		AverageRating: loader.StrToFloat(xs[3]),
		ISBN:          xs[4],
		ISBN13:        xs[5],
		LanguageCode:  xs[6],
		NumPages:      loader.StrToInt(xs[7]),
		Ratings:       loader.StrToInt(xs[8]),
		Reviews:       loader.StrToInt(xs[9]),
	}
}

// Initialize is the method used to populate the in-memory datastore.
// At the beginning, this simply returns a pointer to the struct literal.
// You need to change this to load data from the CSV file
func (b *Books) Initialize() {
	// Measure time
	start := time.Now()
	defer log.Printf("Read operation took %s", time.Since(start))
	// Open csv file
	f, err := os.Open("assets/books.csv")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	// Process file content
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}
	bookdata := []*loader.BookData{}
	for _, line := range lines {
		bd := b.strToBookData(line)
		bookdata = append(bookdata, &bd)
	}
	b.Store = &bookdata
}

// GetAllBooks returns the entire dataset, subjet to the rudimentary limit & skip parameters
func (b *Books) GetAllBooks(limit, skip int) *[]*loader.BookData {
	if limit == 0 || limit > len(*b.Store) {
		limit = len(*b.Store)
	}
	ret := (*b.Store)[skip:limit]
	return &ret
}

// GetAllBooksByAuthor returns the entire dataset of books by the specified author
func (b *Books) GetAllBooksByAuthor(author string) *[]*loader.BookData {
	author = strings.ToLower(author)
	ret := []*loader.BookData{}
	for _, bd := range *b.Store {
		if strings.Contains(strings.ToLower(bd.Authors), author) {
			ret = append(ret, bd)
		}
	}
	return &ret
}

// GetAllBooksByTitle returns the entire dataset of books by the specified title
func (b *Books) GetAllBooksByTitle(title string) *[]*loader.BookData {
	title = strings.ToLower(title)
	ret := []*loader.BookData{}
	for _, bd := range *b.Store {
		if strings.Contains(strings.ToLower(bd.Title), title) {
			ret = append(ret, bd)
		}
	}
	return &ret
}

// GetBookByISBN only returns a first exact match
func (b *Books) GetBookByISBN(isbn string) *loader.BookData {
	isbn = strings.ToLower(isbn)
	for _, bd := range *b.Store {
		if strings.ToLower(bd.ISBN) == isbn {
			print(bd.ISBN)
			return bd
		}
	}
	return nil
}

// RemoveBookByISBN delete book by exact ISBN match
func (b *Books) RemoveBookByISBN(isbn string) bool {
	isbn = strings.ToLower(isbn)
	found := false
	ret := []*loader.BookData{}
	for _, bd := range *b.Store {
		if strings.ToLower(bd.ISBN) != isbn {
			ret = append(ret, bd)
		} else {
			found = true
		}
	}
	b.Store = &ret
	return found
}

// AddBook add new book to the store
func (b *Books) AddBook(book loader.BookData) {
	*b.Store = append(*b.Store, &book)
}
