package request

type ReturnRequest struct {
	BorrowID uint `json:"borrow_id"`
	UserID   uint `json:"user_id"`
	BookID   uint `json:"book_id"`
}
