package request

import "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"

type SearchBook struct {
	Title       string `form:"title"`
	Description string `form:"description"`
	Quantity    uint   `form:"quantity"`
	Cover       string `form:"cover"`
}

func (search *SearchBook) SearchBookToBook() *models.Book {
	return &models.Book{
		Title:       search.Title,
		Description: search.Description,
		Quantity:    search.Quantity,
		Cover:       search.Cover,
	}
}
