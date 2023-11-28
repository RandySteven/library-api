package repository

import (
	"context"
	"fmt"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/query"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// Find implements interfaces.UserRepository.
func (repo *userRepository) Find(ctx context.Context, whereClause []query.WhereClause) ([]models.User, error) {
	var users []models.User
	table := repo.db.WithContext(ctx).Model(&models.User{})
	for _, clause := range whereClause {
		query := fmt.Sprintf("%s %s ?", clause.Field, clause.Condition)
		if len(clause.Value) > 1 {
			table = table.Where(query, "%"+clause.Value+"%")
		}

	}

	err := table.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Save implements interfaces.UserRepository.
func (repo *userRepository) Save(ctx context.Context, user *models.User) (*models.User, error) {
	err := repo.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

var _ interfaces.UserRepository = &userRepository{}
