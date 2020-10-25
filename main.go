package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book model
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author model
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Book slice struct
var books []Book

// get all Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// get Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// create Book
func createBook(w http.ResponseWriter, r *http.Request) {

}

// update Book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// delete Book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// init router
	r := mux.NewRouter()

	books = append(books, Book{ID: "1", Isbn: "643423", Title: "Alchemist", Author: &Author{Firstname: "Paulo", Lastname: "Coelho"}})
	books = append(books, Book{ID: "2", Isbn: "25533", Title: "The subtle art of not giving a f*ck", Author: &Author{Firstname: "Mark", Lastname: "Manson"}})
	books = append(books, Book{ID: "3", Isbn: "123124", Title: "Who moved my cheese", Author: &Author{Firstname: "Spencer", Lastname: "Johnson"}})

	// route endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
