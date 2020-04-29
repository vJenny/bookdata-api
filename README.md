# bookdata-api

This is a test API used for learning.

## Challenge 2  learnings
- files that have the same package declaration will be found, regardless of filename (e.g. routes.go, my new utils.go)
- have to create util functions (for error handling) to use a function that returns multiple values to populate a struct (e.g. strconv)
- need to exit devcontainer to be able to commit to git
- TODO: error not handled if skip > limit
- pay attention to brackets when checking Type: `fmt.Printf("%T\n", rows)` says string but it's an array of arrays ([][]string)
- need to go back to tour for better understanding of pointers (* and &)