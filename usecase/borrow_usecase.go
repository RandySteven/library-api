package usecase

import (
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/enums"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/query"
)

type borrowUseCase struct {
	repo interfaces.BorrowRepository
}

// ReturnBorrowedBook implements interfaces.BorrowUseCase.
func (usecase *borrowUseCase) ReturnBorrowedBookByBorrowId(id uint) (*models.Borrow, error) {
	return usecase.repo.ReturnBookByBorrowId(id)
}

// CreateBorrowRecord implements interfaces.BorrowUseCase.
func (usecase *borrowUseCase) CreateBorrowRecord(borrow *models.Borrow) (*models.Borrow, error) {
	borrow.BorrowStatus = enums.Borrowed
	borrow.BorrowingDate = time.Now()

	return usecase.repo.Save(borrow)
}

// GetAllBorrowsRecord implements interfaces.BorrowUseCase.
func (usecase *borrowUseCase) GetAllBorrowsRecord(whereClauses []query.WhereClause) ([]models.Borrow, error) {
	return usecase.repo.Find(whereClauses)
}

func NewBorrowUseCase(repo interfaces.BorrowRepository) *borrowUseCase {
	return &borrowUseCase{repo: repo}
}

var _ interfaces.BorrowUseCase = &borrowUseCase{}
