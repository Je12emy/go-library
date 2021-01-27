package service

import (
	"library/domain"
	"library/dto"
	"library/errs"
)

// BookService Interface for the Book Service
type BookService interface {
	CreateNewBook(dto.BookRequest) (*dto.BookResponse, *errs.AppError)
	RetrieveBook(dto.BookRequest) (*dto.BookResponse, *errs.AppError)
	// UpdateBook(*dto.BookRequest) (*dto.BookResponse, *errs.AppError)
	// DeleteBook(*dto.BookRequest) (*dto.BookResponse, *errs.AppError)
	RetrieveAllBooks() ([]dto.BookResponse, *errs.AppError)
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

// RetrieveAllBooks Returns all books in the library
func (d DefaultBookService) RetrieveAllBooks() ([]dto.BookResponse, *errs.AppError) {
	books, err := d.repo.FindAll(0, 10)

	if err != nil {
		return nil, err
	}

	var booksResponse []dto.BookResponse
	// for each loop
	for _, b := range books {
		booksResponse = append(booksResponse, b.ToDTO())
	}
	return booksResponse, nil
}

// NewBookService Creates a new Default Book Service implementation
func NewBookService(repo domain.BookRepository) DefaultBookService {
	return DefaultBookService{repo}
}
