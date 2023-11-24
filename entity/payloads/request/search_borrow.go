package request

type SearchBorrow struct {
	UserID uint `form:"user_id"`
	BookID uint `form:"book_id"`
}
