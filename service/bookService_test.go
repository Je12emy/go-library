package service

import (
	realDomain "library/domain"
	"library/dto"
	"library/mocks/domain"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
)

var mockRepo *domain.MockBookRepository
var ctrl gomock.Controller
var service BookService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockRepo = domain.NewMockBookRepository(ctrl)
	service = NewBookService(mockRepo)

	return func() {
		service = nil
		defer ctrl.Finish()
	}
}

func Test_should_new_book_response_when_book_is_saved_successfully(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	req := dto.BookRequest{
		Name:            "My new Book",
		PublicationDate: "14/05/20",
		Genre:           "Terror",
	}

	book := realDomain.Book{
		Name:            req.Name,
		PublicationDate: req.PublicationDate,
		Genre:           req.Genre,
	}
	bookWithID := book
	bookWithID.ID = 1

	mockRepo.EXPECT().Create(book).Return(&bookWithID, nil)

	// Act
	newBook, err := service.CreateNewBook(req)

	// Assert
	if err != nil {
		t.Error("Test failed while creating a new book")
	}

	// TODO: Fix types
	id, _ := strconv.ParseUint(newBook.ID, 10, 32)
	if id != uint64(bookWithID.ID) {
		t.Error("Tested failed while matching new book id")
	}

}