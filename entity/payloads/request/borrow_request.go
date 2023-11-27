package request

type BorrowRequest struct {
	BookID uint `json:"book_id" binding:"required"`
}
