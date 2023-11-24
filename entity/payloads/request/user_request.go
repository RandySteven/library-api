package request

type UserRequest struct {
	Name        string `json:"name" binding:"required,max=35"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required,max=15"`
}
