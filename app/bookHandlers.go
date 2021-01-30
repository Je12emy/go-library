package app

import (
	"encoding/json"
	"library/dto"
	"library/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type BookHandler struct {
	service service.BookService
}

// NewBook Function handler for creating a new book
func (b BookHandler) NewBook(w http.ResponseWriter, r *http.Request) {
	var request dto.BookRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
	} else {
		book, appErr := b.service.CreateNewBook(request)

		if err != nil {
			WriteResponse(w, appErr.Code, appErr.Message)
		} else {
			WriteResponse(w, http.StatusCreated, book)
		}
	}
}

// book/{id}

// FindBook Returns a book in the database
func (b BookHandler) FindBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bookID := vars["book_id"]

	var request dto.BookRequest
	ID, err := strconv.ParseUint(bookID, 10, 32)

	if err != nil {
		WriteResponse(w, http.StatusUnprocessableEntity, "The id you provided is not valid")
	}

	request.ID = uint(ID)
	book, appErr := b.service.RetrieveBook(request)

	if appErr != nil {
		WriteResponse(w, appErr.Code, appErr.Message)
	} else {
		WriteResponse(w, http.StatusFound, book)
	}
}

// UpdateBook Update handler
func (b BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bookID := vars["book_id"]

	var request dto.BookRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	ID, err := strconv.ParseUint(bookID, 10, 32)
	request.ID = uint(ID)

	if err != nil {
		WriteResponse(w, http.StatusUnprocessableEntity, "The request body is invalid")
	} else {
		book, appErr := b.service.UpdateBook(request)

		if appErr != nil {
			WriteResponse(w, appErr.Code, appErr.Message)
		} else {
			WriteResponse(w, http.StatusOK, book)
		}
	}
}
