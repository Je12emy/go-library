package domain

import (
	"library/errs"

	"gorm.io/gorm"
)

type BookRepositoryDB struct {
	dbClient *gorm.DB
}

// todo: check for NOT FOUND error and out of page?

// FindBy Returns a single book by it's id
func (d BookRepositoryDB) FindBy(id int) (*Book, *errs.AppError) {
	var book Book
	result := d.dbClient.First(&book, id)

	if result.RowsAffected == 0 {
		return nil, errs.NewNotFoundError("Book not found")
	}
	if result.Error != nil {
		return nil, errs.NewUnexpectedError("Error while accesing database")
	}
	return &book, nil
}

// FindAll Returns a all the books with pagination, if 0 is passed no limit is imposed.
func (d BookRepositoryDB) FindAll(limit int, offset int) ([]Book, *errs.AppError) {
	books := make([]Book, 0)
	var result *gorm.DB
	if limit == 0 && offset == 0 {
		result = d.dbClient.Find(&books)
	} else {
		result = d.dbClient.Offset(offset).Limit(limit).Find(&books)
	}

	if result.Error != nil {
		return nil, errs.NewUnexpectedError("Error while accesing database")
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
