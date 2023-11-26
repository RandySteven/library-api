package repository

import (
	"fmt"
	"log"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/enums"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/query"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type borrowRepository struct {
	db *gorm.DB
}

// FindBorrowRecordById implements interfaces.BorrowRepository.
func (repo *borrowRepository) FindBorrowRecordById(id uint) (*models.Borrow, error) {
	var borrow models.Borrow
	err := repo.db.Model(&models.Borrow{}).
		Preload("User").
		Preload("Book").
		Where("id = ? ", id).
		Scan(&borrow).
		Error
	if err != nil {
		return nil, err
	}
	return &borrow, nil
}

// ReturnBook implements interfaces.BorrowRepository.
func (repo *borrowRepository) ReturnBookByBorrowId(id uint) (*models.Borrow, error) {
	tx := repo.db.Begin()
	var borrow *models.Borrow

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Model(&models.Borrow{}).
		Where("id = ? ", id).
		Scan(&borrow).Error
	log.Println(borrow)
	if err != nil || borrow == nil {
		return nil, err
	}

	if borrow.BorrowStatus == enums.Returned {
		return nil, apperror.NewErrBorrowStatusAlreadyReturned()
	}

	err = tx.Table("books").
		Where("id = ?", borrow.BookID).
		Update("quantity", gorm.Expr("quantity + ?", 1)).Error
	if err != nil {
		return nil, err
	}

	err = tx.Model(&borrow).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"borrow_status":  enums.Returned,
			"returning_date": time.Now(),
		}).Error
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

	var user *models.User
	err := tx.Model(&models.User{}).
		Where("id = ?", borrow.UserID).
		Scan(&user).Error
	if err != nil || user == nil {
		return nil, apperror.NewErrUserIdNotFound()
	}

	var book *models.Book
	err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).Model(&models.Book{}).
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
