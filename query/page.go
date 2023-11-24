package query

type Page struct {
	Limit  uint   `form:"limit"`
	Offset uint   `form:"offset"`
	Order  string `form:"order"`
}
