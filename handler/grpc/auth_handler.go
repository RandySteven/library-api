package handler_grpc

import (
	"context"
	"log"
	"strings"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/payloads/request"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/pb"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/utils"
)

type AuthHandler struct {
	uscase interfaces.AuthUseCase
	pb.UnimplementedAuthServiceServer
}

func NewAuthHandler(uscase interfaces.AuthUseCase) *AuthHandler {
	return &AuthHandler{uscase: uscase}
}

func (h *AuthHandler) Login(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {

	log.Println("test hit")
	user := &request.UserLoginRequest{
		Email:    strings.TrimSpace(req.Email),
		Password: strings.TrimSpace(req.Password),
	}

	errValidate := utils.Validate(user)
	if errValidate != nil {
		return &pb.AuthResponse{Message: "Login Failed"}, apperror.NewErrBadRequest(errValidate[0])
	}

	token, err := h.uscase.LoginUserByEmail(ctx, user.Email, user.Password)
	if err != nil {
		return &pb.AuthResponse{Message: "Login failed"}, err
	}
	return &pb.AuthResponse{
		Message: "Login success",
		Token:   token,
	}, nil
}
