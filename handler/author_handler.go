package handler

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"github.com/gin-gonic/gin"
)

type AuthorHandler struct {
	usecase interfaces.AuthorUseCase
}

// CreateAuthor implements interfaces.AuthorHandler.
func (*AuthorHandler) CreateAuthor(c *gin.Context) {
	panic("unimplemented")
}

// GetAllAuthors implements interfaces.AuthorHandler.
func (*AuthorHandler) GetAllAuthors(c *gin.Context) {
	panic("unimplemented")
}

func NewAuthorHandler(usecase interfaces.AuthorUseCase) *AuthorHandler {
	return &AuthorHandler{usecase: usecase}
}

var _ interfaces.AuthorHandler = &AuthorHandler{}
