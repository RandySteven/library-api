package handler

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/payloads/request"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/payloads/response"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/enums"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/query"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	usecase interfaces.UserUseCase
}

// CreateUser implements interfaces.UserHandler.
func (handler *UserHandler) CreateUser(c *gin.Context) {
	var request request.UserRequest
	requestId := uuid.NewString()
	ctx := context.WithValue(c.Request.Context(), "request_id", requestId)

	if err := c.ShouldBind(&request); err != nil {
		c.Error(apperror.NewErrBadRequest(err.Error()))
		return
	}

	user := &models.User{
		Name:        request.Name,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
	}

	user, err := handler.usecase.CreateUser(ctx, user)
	if err != nil {
		// resp := response.Response{
		// 	Errors: []string{err.Error()},
		// }
		// c.AbortWithStatusJSON(http.StatusConflict, resp)
		c.Error(err)
		return
	}
	resp := response.Response{
		Message: "Success create user",
		Data:    user,
	}
	c.JSON(http.StatusCreated, resp)
}

// GetAllUsers implements interfaces.UserHandler.
func (handler *UserHandler) GetAllUsers(c *gin.Context) {
	var search request.SearchUser
	requestId := uuid.NewString()
	ctx := context.WithValue(c.Request.Context(), "request_id", requestId)
	c.ShouldBindQuery(&search)

	val := reflect.ValueOf(&search).Elem()
	var whereClauses []query.WhereClause
	for i := 0; i < val.NumField(); i++ {
		whereClause := &query.WhereClause{
			Field:     val.Type().Field(i).Name,
			Value:     fmt.Sprintf("%v", val.Field(i).Interface()),
			Condition: enums.Ilike,
		}
		whereClauses = append(whereClauses, *whereClause)
	}

	users, err := handler.usecase.GetAllUsers(ctx, whereClauses)
	if err != nil {
		resp := response.Response{
			Errors: []string{err.Error()},
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}

	resp := response.Response{
		Message: "Success get all users",
		Data:    users,
	}
	c.JSON(http.StatusOK, resp)
}

func NewUserHandler(usecase interfaces.UserUseCase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

var _ interfaces.UserHandler = &UserHandler{}
