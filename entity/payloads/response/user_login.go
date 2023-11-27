package response

import "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"

type UserLogin struct {
	ID          uint
	Name        string
	Email       string
	PhoneNumber string
}

func NewUserLogin(user *models.User) *UserLogin {
	return &UserLogin{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}
}
