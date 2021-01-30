package service

import (
	realDomain "library/domain"
	"library/dto"
	"library/mocks/domain"
	"testing"
	"time"

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
		ID:              1,
		Name:            "My new Book",
		PublicationDate: "14/05/20",
		Genre:           "Terror",
	}
}

func newBookResponse() []dto.BookResponse {
	return []dto.BookResponse{
		{
			ID:              1,
			Name:            "Book 1",
			PublicationDate: "14/05/20",
			Genre:           "Terror",
		},
		{
			ID:              2,
			Name:            "Book 2",
			PublicationDate: "14/05/20",
			Genre:           "Comedy",
		},
		{
			ID:              3,
			Name:            "Book 3",
			PublicationDate: "14/05/20",
			Genre:           "Parody",
		},
	}
}

func newBookSlice() []realDomain.Book {
	return []realDomain.Book{
		{
			ID:              1,
			Name:            "Book 1",
			PublicationDate: "14/05/20",
			Genre:           "Terror",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
		{
			ID:              2,
			Name:            "Book 2",
			PublicationDate: "14/05/20",
			Genre:           "Comedy",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
		{
			ID:              3,
			Name:            "Book 3",
			PublicationDate: "14/05/20",
			Genre:           "Parody",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
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

	if newBook.ID != bookWithID.ID {
		t.Error("Test failed while matching new book id")
	}
}

func Test_should_find_a_single_book(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	req := NewBookRequest()

	book := realDomain.Book{
		ID:              req.ID,
		Name:            req.Name,
		PublicationDate: req.PublicationDate,
		Genre:           req.Genre,
	}

	mockRepo.EXPECT().FindBy(int(book.ID)).Return(&book, nil)
	b, err := service.RetrieveBook(req)

	// Act
	if err != nil {
		t.Error("Test failed while finding book")
	}

	if b.ID != book.ID {
		t.Error("Test failed while finding book")
	}
}

func Test_should_update_existing_book(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	req := NewBookRequest()

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

	if res.ID != b.ID {
		t.Error("Test failed since returned if does not match")
	}
}

func Test_should_delete_a_single_book(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	req := NewBookRequest()

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

	if res.ID != b.ID {
		t.Error("Test failed since returned if does not match")
	}
}

func Test_should_all_books(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	var pageResponse dto.BookPaginationResponse

	pageResponse.PageInfo.NextPage = "localhost:8000/books?page=1"
	pageResponse.PageInfo.PreviousPage = ""

	pageResponse.Book = newBookResponse()

	mockRepo.EXPECT().FindAll(1).Return(newBookSlice(), nil)
	mockRepo.EXPECT().FindAll(2).Return(newBookSlice(), nil)

	// Act
	_, err := service.RetrieveAllBooks(1)

	// Assert
	if err != nil {
		t.Error("Failed while retrieving all books")
	}
}
