package app

import (
	"encoding/json"
	"library/dto"
	"library/service"
	"net/http"
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
