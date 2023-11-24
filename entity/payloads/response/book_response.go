package response

import (
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"gorm.io/gorm"
)

type BookResponse struct {
	ID          uint            `json:"id"`
	Title       string          `json:"name"`
	Description string          `json:"description"`
	Quantity    uint            `json:"quantity"`
	Cover       string          `json:"cover"`
	Author      *models.Author  `json:"author,omitempty"`
	CreatedAt   *time.Time      `json:"created_at,omitempty"`
	UpdatedAt   *time.Time      `json:"updated_at,omitempty"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func NewBookResponse(book *models.Book, withAuthor bool) *BookResponse {
	bookResponse := &BookResponse{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		Quantity:    book.Quantity,
		Cover:       book.Cover,
		CreatedAt:   &book.CreatedAt,
		UpdatedAt:   &book.UpdatedAt,
		DeletedAt:   &book.DeletedAt,
	}
	if withAuthor {
		bookResponse.Author = &book.Author
		return bookResponse
	}
	return bookResponse
}

// func (response *BookResponse) SetAdditionalInfo(book *models.Book) {
// }
