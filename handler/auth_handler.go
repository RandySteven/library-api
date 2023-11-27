package handler

import (
	"context"
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/payloads/request"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/payloads/response"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthHandler struct {
	usecase interfaces.AuthUseCase
}

func NewAuthHandler(usecase interfaces.AuthUseCase) *AuthHandler {
	return &AuthHandler{usecase: usecase}
}

// LogoutUser handles the route for logging out a user.
func (h *AuthHandler) LogoutUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})
	resp := response.Response{
		Message: "Success to logout",
	}
	c.JSON(http.StatusOK, resp)
}

// LoginUser handles the route for logging in a user.
func (h *AuthHandler) LoginUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	requestId := uuid.NewString()
	ctx := context.WithValue(c.Request.Context(), "request_id", requestId)
	var request request.UserLoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		resp := response.Response{
			Errors: []string{err.Error()},
		}
		utils.ResponseHandler(c.Writer, http.StatusInternalServerError, resp)
		return
	}

	pass, err := utils.HashPassword(request.Password)
	request.Password = pass
	if err != nil {
		resp := response.Response{
			Errors: []string{err.Error()},
		}
		utils.ResponseHandler(c.Writer, http.StatusInternalServerError, resp)
		return
	}

	token, err := h.usecase.LoginUserByEmail(ctx, request.Email, request.Password)
	if err != nil {
		resp := response.Response{
			Errors: []string{err.Error()},
		}
		utils.ResponseHandler(c.Writer, http.StatusInternalServerError, resp)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	resp := response.Response{
		Message: "Success login user",
		Data:    token,
	}
	c.JSON(http.StatusOK, resp)
}

// RegisterUser handles the route for registering a new user.
func (h *AuthHandler) RegisterUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	requestId := uuid.NewString()
	ctx := context.WithValue(c.Request.Context(), "request_id", requestId)

	var register request.UserRequest
	if err := c.ShouldBindJSON(&register); err != nil {
		resp := response.Response{
			Errors: []string{err.Error()},
		}
		utils.ResponseHandler(c.Writer, http.StatusInternalServerError, resp)
		return
	}

	validationErr := utils.Validate(register)
	if validationErr != nil {
		resp := response.Response{
			Errors: validationErr,
		}
		utils.ResponseHandler(c.Writer, http.StatusInternalServerError, resp)
		return
	}

	user := &models.User{
		Name:        register.Name,
		Email:       register.Email,
		Password:    register.Password,
		PhoneNumber: register.PhoneNumber,
	}

	pass, err := utils.HashPassword(user.Password)
	if err != nil {
		resp := response.Response{
			Errors: []string{err.Error()},
		}
		utils.ResponseHandler(c.Writer, http.StatusInternalServerError, resp)
		return
	}
	user.Password = pass
	userStore, err := h.usecase.RegisterUser(ctx, user)
	if err != nil {
		resp := response.Response{
			Errors: []string{err.Error()},
		}
		utils.ResponseHandler(c.Writer, http.StatusInternalServerError, resp)
		return
	}
	resp := response.Response{
		Message: "Success created user",
		Data:    userStore,
	}
	c.JSON(http.StatusOK, resp)
}

var _ interfaces.AuthHandler = &AuthHandler{}
