package repository

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

// RegisterUser implements interfaces.AuthRepository.
func (repo *authRepository) RegisterUser(user *models.User) (*models.User, error) {
	if err := repo.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// LoginUserByEmail implements interfaces.AuthRepository.
func (repo *authRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db}
}

var _ interfaces.AuthRepository = &authRepository{}
