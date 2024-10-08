package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/caslanqa/go-bookstore/pkg/models"
	"github.com/caslanqa/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetBooks()

	res, _ := json.Marshal(newBooks)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)

	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}

	utils.ParseBody(r, CreateBook)
	book := CreateBook.CreateBook()

	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	models.DeleteBook(ID)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("[]"))
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updatedBook := &models.Book{}
	utils.ParseBody(r, updatedBook)
	vars := mux.Vars(r)

	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db := models.GetBookById(ID)
	if updatedBook.Name != "" {
		bookDetails.Name = updatedBook.Name
	}

	if updatedBook.Author != "" {
		bookDetails.Author = updatedBook.Author
	}

	if updatedBook.Publication != "" {
		bookDetails.Publication = updatedBook.Publication
	}

	db.Save(&bookDetails)

	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
