package dto

// BookRequest Request DTO
type BookRequest struct {
	ID              uint   `json:"book_id"`
	Name            string `json:"book_name"`
	PublicationDate string `json:"publication_date"`
	Genre           string `json:"book_genre"`
}
