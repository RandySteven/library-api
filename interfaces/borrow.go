package interfaces

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/query"
	"github.com/gin-gonic/gin"
)

type (
	BorrowRepository interface {
		Save(ctx context.Context, borrow *models.Borrow) (*models.Borrow, error)
		Find(ctx context.Context, whereClauses []query.WhereClause) ([]models.Borrow, error)
		ReturnBookByBorrowId(ctx context.Context, id uint) (*models.Borrow, error)
		FindBorrowRecordById(ctx context.Context, id uint) (*models.Borrow, error)
	}

	BorrowUseCase interface {
		CreateBorrowRecord(ctx context.Context, borrow *models.Borrow) (*models.Borrow, error)
		GetAllBorrowsRecord(ctx context.Context, whereClauses []query.WhereClause) ([]models.Borrow, error)
		ReturnBorrowedBookByBorrowId(ctx context.Context, borrowId uint, userId uint) (*models.Borrow, error)
	}

	BorrowHandler interface {
		CreateBorrowRecord(c *gin.Context)
		GetAllBorrowsRecord(c *gin.Context)
		ReturnBorrowBook(c *gin.Context)
	}
)
