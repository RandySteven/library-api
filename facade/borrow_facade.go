package facade

import "git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/configs"

type BorrowFacade struct {
	repo *configs.Repository
}

func NewBorrowFacade(repo *configs.Repository) *BorrowFacade {
	return &BorrowFacade{repo}
}

func (b *BorrowFacade) BorrowTransaction() {
	tx := b.repo.TrxBegin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

}
