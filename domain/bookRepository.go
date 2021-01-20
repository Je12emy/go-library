package domain

import (
	"gorm.io/gorm"
)

type BookRepositoryDB struct {
	dbClient *gorm.DB
}

// FindBy Returns a single book by it's id
func (d BookRepositoryDB) FindBy(id int) (*Book, error) {
	var book Book
	result := d.dbClient.First(&book, id)

	if result.Error != nil {
		return nil, result.Error
	}
	// if result.RowsAffected == 0 {
	// 	fmt.Println("book not found")
	// }
	return &book, nil
}

// NewBookRepositoryDB Returns a new instance of BookRepository takes a gorm db client
func NewBookRepositoryDB(dbClient *gorm.DB) BookRepositoryDB {
	return BookRepositoryDB{dbClient}
}
