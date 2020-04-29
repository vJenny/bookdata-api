package datastore

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/matt-FFFFFF/bookdata-api/loader"
)

// Books is the memory-backed datastore used by the API
// It contains a single field 'Store', which is (a pointer to) a slice of loader.BookData struct pointers
type Books struct {
	Store *[]*loader.BookData `json:"store"`
}

// Initialize is the method used to populate the in-memory datastore.
// At the beginning, this simply returns a pointer to the struct literal.
// You need to change this to load data from the CSV file
func (b *Books) Initialize() {
	// start timer and defer finish
	start := time.Now()
	defer log.Printf("Read operation too %s", time.Since(start))

	file, err := os.Open("./assets/books.csv")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	rows, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	bd := []*loader.BookData{}
	for _, row := range rows {
		// strconv returns two values, so have to create a util funtion.
		// put this in a utils file in the same package (same package declaration),
		// which means it is found - no need to point at it.
		item := loader.BookData{
			BookID:        row[0],
			Title:         row[1],
			Authors:       row[2],
			AverageRating: utilStrToFloat(row[3]),
			ISBN:          row[4],
			ISBN13:        row[5],
			LanguageCode:  row[6],
			NumPages:      utilStrToInt(row[7]),
			Ratings:       utilStrToInt(row[8]),
			Reviews:       utilStrToInt(row[9]),
		}
		bd = append(bd, &item)
	}

	// diagnostics
	fmt.Println(len(rows))
	fmt.Printf("%T\n", rows)
	fmt.Println(rows[0])
	fmt.Println(rows[1])
	fmt.Println(rows[1][1])

	//b.Store = &loader.BooksLiteral
	b.Store = &bd
}

// GetAllBooks returns the entire dataset, subjet to the rudimentary limit & skip parameters
func (b *Books) GetAllBooks(limit, skip int) *[]*loader.BookData {
	if limit == 0 || limit > len(*b.Store) {
		limit = len(*b.Store)
	}
	ret := (*b.Store)[skip:limit]
	return &ret
}
