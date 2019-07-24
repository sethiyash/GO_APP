package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"math/rand"
)

//Book struct (Model)
type Book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

// int book var as slice book struct (slice means arrau=y of variable length)
var books []Book 

// Get all the books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(books)
}

// Get all the books
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for _,item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Get all the books
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa( rand.Intn(100) )
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Get all the books
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index],books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Get all the books
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index],books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main()  {
	// Init Router
	r := mux.NewRouter()

	//Mock Data Initializing books
	books = append(books, Book{ID: "1", Title: "Python Book", Author: &Author{Firstname: "Mark", Lastname: "Juckerberg"}})
	books = append(books, Book{ID: "2", Title: "GO Book", Author: &Author{Firstname: "John", Lastname: "Marker"}})


	//Route Handelers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
