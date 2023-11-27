package usecase

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/configs"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
)

type Usecase struct {
	interfaces.BookUseCase
	interfaces.AuthorUseCase
	interfaces.UserUseCase
	interfaces.BorrowUseCase
}

func NewUsecase(repo configs.Repository) *Usecase {
	return &Usecase{
		BookUseCase:   NewBookUseCase(repo.BookRepository, repo.AuthorRepository),
		AuthorUseCase: NewAuthorUseCase(repo.AuthorRepository),
		UserUseCase:   NewUserUseCase(repo.UserRepository),
		BorrowUseCase: NewBorrowUseCase(repo.BorrowRepository),
	}
}
