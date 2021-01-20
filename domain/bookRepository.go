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

// FindAll Returns a all the books with pagination, if 0 is passed no limit is imposed.
func (d BookRepositoryDB) FindAll(limit int, offset int) ([]Book, error) {
	books := make([]Book, 0)
	var result *gorm.DB
	if limit == 0 && offset == 0 {
		result = d.dbClient.Find(&books)
	} else {
		result = d.dbClient.Offset(offset).Limit(limit).Find(&books)
	}

	if result.Error != nil {
		return nil, result.Error
	}
	// if result.RowsAffected == 0 {
	// 	fmt.Println("book not found")
	// }
	return books, nil
}

// NewBookRepositoryDB Returns a new instance of BookRepository takes a gorm db client
func NewBookRepositoryDB(dbClient *gorm.DB) BookRepositoryDB {
	return BookRepositoryDB{dbClient}
}
