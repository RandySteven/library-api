package usecase

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/query"
)

type userUseCase struct {
	repo interfaces.UserRepository
}

// CreateUser implements interfaces.UserUseCase.
func (usecase *userUseCase) CreateUser(user *models.User) (*models.User, error) {
	return usecase.repo.Save(user)
}

// GetAllUsers implements interfaces.UserUseCase.
func (usecase *userUseCase) GetAllUsers(whereClause []query.WhereClause) ([]models.User, error) {
	return usecase.repo.Find(whereClause)
}

func NewUserUseCase(repo interfaces.UserRepository) *userUseCase {
	return &userUseCase{repo: repo}
}

var _ interfaces.UserUseCase = &userUseCase{}
