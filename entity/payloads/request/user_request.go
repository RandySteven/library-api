package request

type UserRequest struct {
	Name        string `json:"name" binding:"required,max=35"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required,max=15"`
	Password    string `json:"password" binding:"required,min=8"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required" validate:"required|email"`
	Password string `json:"password" binding:"required,min=8" validate:"required|min=8|max=16"`
}
