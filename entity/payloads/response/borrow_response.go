package response

import "time"

type BorrowResponse struct {
	ID           uint      `json:"id"`
	BorrowStatus string    `json:"borrow_status"`
	UserID       uint      `json:"user_id"`
	UserName     string    `json:"user_name"`
	BookID       uint      `json:"book_id"`
	BookName     string    `json:"book_name"`
	BookQuantity uint      `json:"book_quantity"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
