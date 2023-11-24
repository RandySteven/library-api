package request

type SearchUser struct {
	Name        string `form:"name"`
	Email       string `form:"email"`
	PhoneNumber string `form:"phone_number"`
}
