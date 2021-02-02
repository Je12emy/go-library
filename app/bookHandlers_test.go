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

	router.HandleFunc("/books/{book_id:[0-9]+}", bh.FindBook).Methods(http.MethodGet)
	router.HandleFunc("/books", bh.NewBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{book_id:[0-9]+}", bh.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{book_id:[0-9]+}", bh.DeleteBook).Methods(http.MethodDelete)

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
		ID: 1,
	}

	mockService.EXPECT().DeleteBook(bookRequest).Return(&bookResponse, nil)

	request, _ := http.NewRequest(http.MethodDelete, "/books/1", nil)
	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Error while validating status code")
	}
}

func Test_should_return_the_updated_book_with_status_200_ok(t *testing.T) {
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
		ID:              1,
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
	mockService.EXPECT().UpdateBook(bookRequest).Return(&bookResponse, nil)
	request, _ := http.NewRequest(http.MethodPut, "/books/1", bytes.NewBuffer(bytesRepresentation))

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Error while validating status code")
	}
}

func Test_should_return_the_deleted_book_with_status_200_ok(t *testing.T) {

}
