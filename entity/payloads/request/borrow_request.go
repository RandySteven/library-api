package request

type BorrowRequest struct {
	UserID uint `json:"user_id"`
	BookID uint `json:"book_id"`
}
