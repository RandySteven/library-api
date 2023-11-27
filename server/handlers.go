package server

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/configs"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/usecase"
)

type (
	Handlers struct {
		BookHandler   interfaces.BookHandler
		UserHandler   interfaces.UserHandler
		BorrowHandler interfaces.BorrowHandler
		AuthHandler   interfaces.AuthHandler
	}
)

func NewHandlers(repo configs.Repository) (*Handlers, error) {
	usecase := usecase.NewUsecase(repo)

	return &Handlers{
		BookHandler:   handler.NewBookHandler(usecase.BookUseCase),
		UserHandler:   handler.NewUserHandler(usecase.UserUseCase),
		BorrowHandler: handler.NewBorrowHandler(usecase.BorrowUseCase),
		AuthHandler:   handler.NewAuthHandler(usecase.AuthUseCase),
	}, nil
}
