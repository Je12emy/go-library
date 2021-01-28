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

func NewBookRequest() dto.BookRequest {
	return dto.BookRequest{
		Name:            "My new Book",
		PublicationDate: "14/05/20",
		Genre:           "Terror",
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
		t.Error("Test failed while matching new book id")
	}
}

func Test_should_find_a_single_book(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	req := NewBookRequest()
	req.ID = 1

	book := realDomain.Book{
		Name:            req.Name,
		PublicationDate: req.PublicationDate,
		Genre:           req.Genre,
	}

	bookWithID := book
	bookWithID.ID = 1

	mockRepo.EXPECT().FindBy(int(bookWithID.ID)).Return(&bookWithID, nil)
	b, err := service.RetrieveBook(req)

	// Act
	if err != nil {
		t.Error("Test failed while finding book")
	}
	// TODO Fix types
	id, _ := strconv.ParseUint(b.ID, 10, 32)
	if id != uint64(bookWithID.ID) {
		t.Error("Test failed while finding book")
	}
}

func Test_should_update_existing_book(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	req := NewBookRequest()
	req.ID = 1

	b := realDomain.Book{
		ID:              req.ID,
		Name:            req.Name,
		Genre:           req.Genre,
		PublicationDate: req.PublicationDate,
	}

	mockRepo.EXPECT().Update(b).Return(&b, nil)

	// Act
	res, err := service.UpdateBook(req)

	// Assert
	if err != nil {
		t.Error("Test failed while updating book")
	}

	id, _ := strconv.ParseUint(res.ID, 10, 32)
	if id != uint64(b.ID) {
		t.Error("Test failed since returned if does not match")
	}
}

func Test_should_delete_a_single_book(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	req := NewBookRequest()
	req.ID = 1

	b := realDomain.Book{
		ID: req.ID,
	}

	mockRepo.EXPECT().Delete(b).Return(&b, nil)

	// Act
	res, err := service.DeleteBook(req)

	// Assert

	if err != nil {
		t.Error("Failed while deleting book")
	}

	id, _ := strconv.ParseUint(res.ID, 10, 32)
	if id != uint64(b.ID) {
		t.Error("Test failed since returned if does not match")
	}
}
