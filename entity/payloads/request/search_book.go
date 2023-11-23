package request

type SearchBook struct {
	Title       string `form:"title"`
	Description string `form:"description"`
	Quantity    uint   `form:"quantity"`
	Cover       string `form:"cover"`
}
