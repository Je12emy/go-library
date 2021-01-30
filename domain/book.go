package domain

import (
	"library/dto"
	"library/errs"
	"time"

	"gorm.io/gorm"
)

// Book Domain object for books
type Book struct {
	ID              uint `gorm:"primaryKey"`
	Name            string
	PublicationDate string
	Genre           string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

// BookRepository Interface for CRUD operations
//go:generate mockgen -destination=../mocks/domain/mockBookRepository.go -package=domain library/domain BookRepository
type BookRepository interface {
	Create(Book) (*Book, *errs.AppError)
	FindBy(bookId int) (*Book, *errs.AppError)
	FindAll(page int) ([]Book, *errs.AppError)
	Update(Book) (*Book, *errs.AppError)
	Delete(Book) (*Book, *errs.AppError)
}

// ToDTO Tranforms the book domain object into DTO
func (b Book) ToDTO() dto.BookResponse {
	return dto.BookResponse{
		ID:              b.ID,
		Name:            b.Name,
		PublicationDate: b.PublicationDate,
		Genre:           b.Genre,
	}
}
