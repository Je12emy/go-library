package app

import (
	"bytes"
	"encoding/json"
	"library/dto"
	"library/mocks/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

var bh BookHandler
var router *mux.Router
var mockService *service.MockBookService

func setup(t *testing.T) func() {
	ctlr := gomock.NewController(t)
	mockService = service.NewMockBookService(ctlr)
	bh = BookHandler{mockService}
	router = mux.NewRouter()
	router.HandleFunc("/books/{book_id:[0-9]+}", bh.FindBook)
	router.HandleFunc("/books", bh.NewBook)

	return func() {
		router = nil
		defer ctlr.Finish()
	}
}

func Test_should_return_a_book_with_status_code_200(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	dummyBook := dto.BookRequest{
		ID: 1,
	}

	response := dto.BookResponse{
		ID:              1,
		Name:            "Book 1",
		PublicationDate: "14/05/2020",
		Genre:           "Comedy",
	}

	mockService.EXPECT().RetrieveBook(dummyBook).Return(&response, nil)
	request, _ := http.NewRequest(http.MethodGet, "/books/1", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	// Assert

	if recorder.Code != http.StatusFound {
		t.Error("Failed while testing status code")
	}
}

func Test_should_return_a_created_book_with_status_200_ok(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	bookResponse := dto.BookResponse{
		ID:              1,
		Name:            "Book 1",
		PublicationDate: "14/05/2020",
		Genre:           "Comedy",
	}

	bookRequest := dto.BookRequest{
		Name:            "Book 1",
		PublicationDate: "14/05/2020",
		Genre:           "Comedy",
	}

	message := map[string]interface{}{
		"book_name":        "Book 1",
		"publication_date": "14/05/2020",
		"book_genre":       "Comedy",
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		t.Error("Error while creating request body")
	}
	mockService.EXPECT().CreateNewBook(bookRequest).Return(&bookResponse, nil)

	request, _ := http.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(bytesRepresentation))
	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	// Assert
	if recorder.Code != http.StatusCreated {
		t.Error("Error while validating status code")
	}
}
