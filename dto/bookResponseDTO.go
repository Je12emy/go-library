package dto

// BookRequest Response DTO
type BookResponse struct {
	ID              uint   `json:"book_id"`
	Name            string `json:"book_name"`
	PublicationDate string `json:"publication_date"`
	Genre           string `json:"book_genre"`
}

type PaginationInfo struct {
	NextPage     string `json:"next"`
	PreviousPage string `json:"prev"`
}

type BookPaginationResponse struct {
	PageInfo PaginationInfo `json:"page"`
	Book     []BookResponse `json:"books"`
}
