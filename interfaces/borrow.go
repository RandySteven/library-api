package interfaces

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/query"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	BorrowRepository interface {
		Save(borrow *models.Borrow, tx *gorm.DB) (*models.Borrow, error)
		Find(whereClauses []query.WhereClause) ([]models.Borrow, error)
		ReturnBookByBorrowId(id uint) (*models.Borrow, error)
		GetBorrowTx() *gorm.DB
	}

	BorrowUseCase interface {
		CreateBorrowRecord(borrow *models.Borrow) (*models.Borrow, error)
		GetAllBorrowsRecord(whereClauses []query.WhereClause) ([]models.Borrow, error)
		ReturnBorrowedBookByBorrowId(id uint) (*models.Borrow, error)
	}

	BorrowHandler interface {
		CreateBorrowRecord(c *gin.Context)
		GetAllBorrowsRecord(c *gin.Context)
		ReturnBorrowBook(c *gin.Context)
	}
)
