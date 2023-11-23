package request

type BookRequest struct {
	Title       string `json:"title" binding:"required,max=35"`
	Description string `json:"description" binding:"required"`
	Quantity    uint   `json:"quantity" binding:"required,min=-1"`
	Cover       string `json:"cover"`
}
