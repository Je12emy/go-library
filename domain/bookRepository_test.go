package domain_test

import (
	"fmt"
	"library/app"
	"library/domain"
	"testing"

	"gorm.io/gorm"
)

var db *gorm.DB

func Test_should_return_a_single_book_by_its_id(t *testing.T) {
	// Arrange
	setup()
	repo := domain.NewBookRepositoryDB(db)

	db.Create(&domain.Book{
		ID:              1,
		Name:            "Elon Musk",
		PublicationDate: "20/20/15",
		Genre:           "Bibliography",
	})

	// Act
	book, err := repo.FindBy(1)
	fmt.Println(book)
	// Assert
	if err != nil {
		t.Error("Failed retrieving book")
	}
}

func setup() {
	db = app.GetTestDBClient()
}
