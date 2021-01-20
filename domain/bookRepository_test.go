package domain_test

import (
	"fmt"
	"library/app"
	"library/domain"
	"testing"

	"gorm.io/gorm"
)

var db *gorm.DB
var repo domain.BookRepositoryDB

func setup() {
	db = app.GetTestDBClient()
	repo = domain.NewBookRepositoryDB(db)
	createBooks()
}

func createBooks() {
	db.Create(&domain.Book{
		ID:              1,
		Name:            "Elon Musk",
		PublicationDate: "20/20/15",
		Genre:           "Bibliography",
	})
	db.Create(&domain.Book{
		ID:              2,
		Name:            "The Master Algorithm",
		PublicationDate: "20/20/15",
		Genre:           "Science & Technology",
	})
	db.Create(&domain.Book{
		ID:              3,
		Name:            "Test Book 3",
		PublicationDate: "20/20/15",
		Genre:           "Science & Technology",
	})
	db.Create(&domain.Book{
		ID:              4,
		Name:            "Test Book 4",
		PublicationDate: "20/20/15",
		Genre:           "Science & Technology",
	})
}

func Test_should_return_a_single_book_by_its_id(t *testing.T) {
	// Arrange
	setup()

	// Act
	book, err := repo.FindBy(1)
	fmt.Println(book)
	// Assert
	if err != nil {
		t.Error("Failed retrieving book")
	}
}

func Test_should_fail_by_not_finding_book_by_its_id(t *testing.T) {
	// Arrange
	setup()

	// Act
	_, err := repo.FindBy(500)
	// Assert
	if err == nil {
		t.Error("Failed while retrieving a non existing book")
	}
}

func Test_should_return_all_books(t *testing.T) {
	setup()

	books, err := repo.FindAll(0, 0)

	if err != nil {
		t.Error("Failed while retrieving all books")
	}

	if len(books) < 1 {
		t.Error("Filed while retrieving the list of all books")
	}
}

func Test_should_paginate_books(t *testing.T) {
	setup()

	booksPage1, err := repo.FindAll(4, 0)

	if err != nil {
		t.Error("Failed while retrieving page 1 of books")
	}

	booksPage2, err := repo.FindAll(4, 2)

	if err != nil {
		t.Error("Failed while retrieving page 2 of books")
	}

	if booksPage1[0].Name == booksPage2[0].Name {
		t.Error("Failed while paginating results")
	}
}

func Test_should_insert_a_new_book(t *testing.T) {
	setup()
	b := domain.Book{
		ID:              999,
		Name:            "Test Book",
		PublicationDate: "12/12/2020",
		Genre:           "Mistery",
	}
	_, err := repo.Create(b)
	if err != nil {
		t.Error("Error while creating a new book")
	}
}
