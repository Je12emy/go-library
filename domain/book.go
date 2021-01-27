package domain

import (
	"library/dto"
	"library/errs"
	"strconv"
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
type BookRepository interface {
	Save(Book) ([]Book, *errs.AppError)
	FindBy(bookId string) (*Book, *errs.AppError)
	FindAll() ([]Book, *errs.AppError)
}

// ToDTO Tranforms the book domain object into DTO
func (b Book) ToDTO() dto.BookResponse {
	return dto.BookResponse{
		ID:              strconv.FormatUint(uint64(b.ID), 10),
		Name:            b.Name,
		PublicationDate: b.PublicationDate,
		Genre:           b.Genre,
	}
}
