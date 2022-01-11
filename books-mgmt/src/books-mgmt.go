package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"log"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func sequence(initValue int) func() int {
	i := initValue

	return func() int {
		i++
		return i
	}
}

// Book represents a single book
type Book struct {
	ID          int
	Title       string
	Author      string
	ISBN        string
	Description string
	Price       float64
}

var nextID = sequence(0)

var books = map[int]Book{
	1: {
		ID:          nextID(),
		Title:       "The C Book",
		Author:      "Dennis Ritchie",
		ISBN:        "",
		Description: "",
		Price:       0,
	},
	2: {
		ID:          nextID(),
		Title:       "C++",
		Author:      "Bjarne Stroustrop",
		ISBN:        "",
		Description: "",
		Price:       0,
	},
}

// Get /
func homeHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(map[string]string{"data": "Books Mgmt"})
}

// GET /books/{id}
func bookShowHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	bookID, _ := strconv.Atoi(vars["id"])
	log.Println("Getting bookID: ", bookID)
	book := books[bookID]

	json.NewEncoder(w).Encode(book)
}

// GET /books
func booksIndexHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func main() {
	port := "9000"
	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler).Methods("GET")

	r.HandleFunc("/books", booksIndexHandler).Methods("GET")
	r.HandleFunc("/books/{id}", bookShowHandler).Methods("GET")

	log.Println("Server is running on port: ", port)

	http.ListenAndServe(":"+port, handlers.LoggingHandler(os.Stdout, r))
}
