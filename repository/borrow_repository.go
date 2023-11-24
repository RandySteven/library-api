package repository

import (
	"fmt"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/enums"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/query"
	"gorm.io/gorm"
)

type borrowRepository struct {
	db *gorm.DB
}

// ReturnBook implements interfaces.BorrowRepository.
func (repo *borrowRepository) ReturnBook(borrow *models.Borrow) (*models.Borrow, error) {
	tx := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Model(&models.Borrow{}).
		Where("id = ? ", borrow.ID).
		Scan(&borrow).Error
	if err != nil || borrow != nil {
		return nil, err
	}

	err = tx.Model(&models.Book{}).
		Where("id = ?", borrow.BookID).
		Update("quantity", "quantity + 1").Error
	if err != nil {
		return nil, err
	}

	err = tx.Model(&models.Borrow{}).
		Where("id = ?", borrow.ID).
		Update("borrow_status", enums.Returned).Error
	if err != nil {
		return nil, err
	}

	tx.Commit()

	return borrow, nil
}

// Find implements interfaces.BorrowRepository.
func (repo *borrowRepository) Find(whereClauses []query.WhereClause) ([]models.Borrow, error) {
	var borrows []models.Borrow
	table := repo.db.Model(&models.Borrow{}).Preload("User").Preload("Book")
	for _, clause := range whereClauses {
		query := fmt.Sprintf("%s %s ?", clause.Field, clause.Condition)
		if len(clause.Value) > 1 {
			table = table.Where(query, "%"+clause.Value+"%")
		}

	}

	err := table.Find(&borrows).Error
	if err != nil {
		return nil, err
	}
	return borrows, nil

}

// Save implements interfaces.BorrowRepository.
func (repo *borrowRepository) Save(borrow *models.Borrow) (*models.Borrow, error) {
	tx := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var book *models.Book
	err := tx.Model(&models.Book{}).
		Where("id = ? ", borrow.BookID).
		Scan(&book).Error
	if err != nil || book == nil {
		return nil, apperror.NewErrBookIdNotFound()
	}

	if book.Quantity == 0 {
		return nil, apperror.NewErrBookQuantityZero()
	}

	err = tx.Table("books").
		Where("id = ?", book.ID).
		Update("quantity", gorm.Expr("quantity - ?", 1)).
		Error

	if err != nil {
		return nil, err
	}

	err = tx.Create(&borrow).Error
	if err != nil {
		return nil, err
	}
	tx.Commit()
	return borrow, nil
}

func NewBorrowRepository(db *gorm.DB) *borrowRepository {
	return &borrowRepository{db}
}

var _ interfaces.BorrowRepository = &borrowRepository{}
