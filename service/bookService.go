package service

import (
	"fmt"
	"library/domain"
	"library/dto"
	"library/errs"
	"os"
)

// BookService Interface for the Book Service
//go:generate mockgen -destination=../mocks/service/mockBookService.go -package=service library/service BookService
type BookService interface {
	CreateNewBook(dto.BookRequest) (*dto.BookResponse, *errs.AppError)
	RetrieveBook(dto.BookRequest) (*dto.BookResponse, *errs.AppError)
	UpdateBook(dto.BookRequest) (*dto.BookResponse, *errs.AppError)
	DeleteBook(dto.BookRequest) (*dto.BookResponse, *errs.AppError)
	RetrieveAllBooks(page int) (*dto.BookPaginationResponse, *errs.AppError)
}

// DefaultBookService Struct for the default book service
type DefaultBookService struct {
	repo domain.BookRepository
}

// CreateNewBook Creates a new Book from a request
func (d DefaultBookService) CreateNewBook(req dto.BookRequest) (*dto.BookResponse, *errs.AppError) {
	book := domain.Book{
		Name:            req.Name,
		PublicationDate: req.PublicationDate,
		Genre:           req.Genre,
	}

	b, err := d.repo.Create(book)
	if err != nil {
		return nil, err
	}

	response := b.ToDTO()
	return &response, nil

}

// RetrieveBook Returns a book by it's id
func (d DefaultBookService) RetrieveBook(req dto.BookRequest) (*dto.BookResponse, *errs.AppError) {

	b, err := d.repo.FindBy(int(req.ID))

	if err != nil {
		return nil, err
	}

	response := b.ToDTO()
	return &response, nil
}

func hasNextPage(books []domain.Book) bool {
	return len(books) == 0
}

// RetrieveAllBooks Returns all books in the library
func (d DefaultBookService) RetrieveAllBooks(page int) (*dto.BookPaginationResponse, *errs.AppError) {

	// Get env server address
	serverAddress := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")

	var response dto.BookPaginationResponse
	var pageInfo dto.PaginationInfo

	books, err := d.repo.FindAll(page)

	if err != nil {
		return nil, err
	}

	if page > 1 {
		pageInfo.PreviousPage = fmt.Sprintf("%v:%v/books?page=%v", serverAddress, serverPort, page-1)
	} else {
		pageInfo.PreviousPage = ""
	}

	nextBooksPage, err := d.repo.FindAll(page + 1)

	if hasNextPage(nextBooksPage) {
		pageInfo.NextPage = ""
	} else {
		pageInfo.NextPage = fmt.Sprintf("%v:%v/books?page=%v", serverAddress, serverPort, page+1)
	}

	response.PageInfo = pageInfo

	var booksResponse []dto.BookResponse

	// for each loop
	for _, b := range books {
		booksResponse = append(booksResponse, b.ToDTO())
	}

	response.Book = booksResponse

	return &response, nil
}

// UpdateBook Updates a book
func (d DefaultBookService) UpdateBook(req dto.BookRequest) (*dto.BookResponse, *errs.AppError) {
	b := domain.Book{
		ID:              req.ID,
		Name:            req.Name,
		PublicationDate: req.PublicationDate,
		Genre:           req.Genre,
	}

	book, err := d.repo.Update(b)

	if err != nil {
		return nil, err
	}
	res := book.ToDTO()
	return &res, nil
}

// DeleteBook Delete a single book
func (d DefaultBookService) DeleteBook(req dto.BookRequest) (*dto.BookResponse, *errs.AppError) {
	b := domain.Book{
		ID: req.ID,
	}

	book, err := d.repo.Delete(b)
	if err != nil {
		return nil, err
	}

	res := book.ToDTO()
	return &res, nil

}

// NewBookService Creates a new Default Book Service implementation
func NewBookService(repo domain.BookRepository) DefaultBookService {
	return DefaultBookService{repo}
}
