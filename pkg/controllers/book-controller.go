package controllers

import (
	"encoding/json"
	"go-mysql/pkg/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetBooks - Get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}
// GetBook - Get a book by ID
func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid book ID")
		return
	}
	book, _ := models.GetBookById(id)
	json.NewEncoder(w).Encode(book)
}

// CreateBook - Create a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid request payload")
		return
	}
	book.CreateBook()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// DeleteBook - Delete a book by ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid book ID")
		return
	}
	book := models.DeleteBook(id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

// UpdateBook - Update a book by ID
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updatedBook = &models.Book{}
	vars := mux.Vars(r)
	bookId := vars["id"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid book ID")
		return
	}
	bookDetails, db := models.GetBookById(id)

	if updatedBook.Title != "" {
		bookDetails.Title = updatedBook.Title
	}
	if updatedBook.Author != "" {
		bookDetails.Author = updatedBook.Author
	}
	if updatedBook.Year != "" {
		bookDetails.Year = updatedBook.Year
	}
	if updatedBook.ISBN != "" {
		bookDetails.ISBN = updatedBook.ISBN
	}
	if updatedBook.Publication != "" {
		bookDetails.Publication = updatedBook.Publication
	}
	db.Save(&bookDetails)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bookDetails)
}