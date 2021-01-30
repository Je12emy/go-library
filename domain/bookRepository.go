package domain

import (
	"library/errs"

	"gorm.io/gorm"
)

type BookRepositoryDB struct {
	dbClient *gorm.DB
}

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

// TODO: Implement pagination by default

// FindAll Returns a all the books with pagination, if 0 is passed no limit is imposed.
func (d BookRepositoryDB) FindAll(page int) ([]Book, *errs.AppError) {

	pagesize := 25

	if page == 0 {
		page = 1
	}

	books := make([]Book, 0)

	offset := (page - 1) * pagesize

	result := d.dbClient.Offset(offset).Limit(pagesize).Find(&books)

	if result.Error != nil {
		return nil, errs.NewUnexpectedError("Error while accesing database")
	}
	// if result.RowsAffected == 0 {
	// 	fmt.Println("book not found")
	// }
	return books, nil
}

// Create Inserts a new book into the database
func (d BookRepositoryDB) Create(book Book) (*Book, *errs.AppError) {
	result := d.dbClient.Create(&book)
	if result.Error != nil {
		return nil, errs.NewUnexpectedError("Error while creating a new book:" + result.Error.Error())
	}
	return &book, nil
}

// Update Updates a book, first it attempts to find the book by it's id and then updates
func (d BookRepositoryDB) Update(book Book) (*Book, *errs.AppError) {
	var result *gorm.DB
	var b Book
	result = d.dbClient.First(&b, book.ID)
	if result.RowsAffected == 0 {
		return nil, errs.NewNotFoundError("Book not found: " + result.Error.Error())
	}

	result = d.dbClient.Save(&book)
	if result.RowsAffected == 0 {
		return nil, errs.NewUnexpectedError("Error while updating book: " + result.Error.Error())
	}

	return &book, nil
}

func (d BookRepositoryDB) Delete(book Book) (*Book, *errs.AppError) {
	var result *gorm.DB
	var b Book

	result = d.dbClient.First(&b, book.ID)
	if result.RowsAffected == 0 {
		return nil, errs.NewNotFoundError("Book not found: " + result.Error.Error())
	}

	result = d.dbClient.Delete(&b)
	if result.RowsAffected == 0 {
		return nil, errs.NewUnexpectedError("Unexpected error while deleting book: " + result.Error.Error())
	}

	return &b, nil
}

// NewBookRepositoryDB Returns a new instance of BookRepository takes a gorm db client
func NewBookRepositoryDB(dbClient *gorm.DB) BookRepositoryDB {
	return BookRepositoryDB{dbClient}
}
