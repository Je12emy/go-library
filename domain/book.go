package domain

import (
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
type BookRepository interface {
	Save(Book) ([]Book, *errs.AppError)
	FindBy(bookId string) (*Book, *errs.AppError)
	FindAll() ([]Book, *errs.AppError)
}
