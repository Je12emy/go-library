package domain

import (
	"time"

	"gorm.io/gorm"
)

// todo: Implement auto-increment keys

type Book struct {
	ID              uint `gorm:"primaryKey"`
	Name            string
	PublicationDate string
	Genre           string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

type BookRepository interface {
	Save(Book) ([]Book, error)
	FindBy(bookId string) (*Book, error)
}
