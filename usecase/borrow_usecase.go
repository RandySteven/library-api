package usecase

import (
	"context"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/enums"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/query"
)

type borrowUseCase struct {
	repo interfaces.BorrowRepository
}

// ReturnBorrowedBook implements interfaces.BorrowUseCase.
func (usecase *borrowUseCase) ReturnBorrowedBookByBorrowId(ctx context.Context, borrowId uint, userId uint) (*models.Borrow, error) {
	borrow, err := usecase.repo.FindBorrowRecordById(ctx, borrowId)
	if err != nil {
		return nil, apperror.NewErrBorrowRecordNotFound().Err
	}
	if borrow.UserID != userId {
		return nil, apperror.NewErrPermissionDenied().Err
	}
	return usecase.repo.ReturnBookByBorrowId(ctx, borrowId)
}

// CreateBorrowRecord implements interfaces.BorrowUseCase.
func (usecase *borrowUseCase) CreateBorrowRecord(ctx context.Context, borrow *models.Borrow) (*models.Borrow, error) {
	borrow.BorrowStatus = enums.Borrowed
	borrow.BorrowingDate = time.Now()

	borrow, err := usecase.repo.Save(ctx, borrow)
	if err != nil {
		return nil, err
	}

	borrowRecord, err := usecase.repo.FindBorrowRecordById(ctx, borrow.ID)
	if err != nil {
		return nil, err
	}

	return borrowRecord, nil
}

// GetAllBorrowsRecord implements interfaces.BorrowUseCase.
func (usecase *borrowUseCase) GetAllBorrowsRecord(ctx context.Context, whereClauses []query.WhereClause) ([]models.Borrow, error) {
	return usecase.repo.Find(ctx, whereClauses)
}

func NewBorrowUseCase(repo interfaces.BorrowRepository) *borrowUseCase {
	return &borrowUseCase{repo: repo}
}

var _ interfaces.BorrowUseCase = &borrowUseCase{}
