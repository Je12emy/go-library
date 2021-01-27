package dto

// BookResponse Request DTO
type BookResponse struct {
	ID              string `json:"book_id"`
	Name            string `json:"book_name"`
	PublicationDate string `json:"publication_date"`
	Genre           string `json:"book_genre"`
}
