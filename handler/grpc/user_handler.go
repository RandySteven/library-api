package handler_grpc

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/pb"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	usecase interfaces.UserUseCase
}

func NewUserHandler(usecase interfaces.UserUseCase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (h *UserHandler) GetAllUsers(ctx context.Context, empty *pb.Empty) (*pb.UserResponsesDTO, error) {

	users, err := h.usecase.GetAllUsers(ctx, nil)
	if err != nil {
		return nil, err
	}

	userResponses := []*pb.UserResponse{}
	for _, user := range users {
		userRes := &pb.UserResponse{
			Id:          uint32(user.ID),
			Name:        user.Name,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			CreatedAt:   user.CreatedAt.GoString(),
			UpdatedAt:   user.UpdatedAt.GoString(),
		}
		userResponses = append(userResponses, userRes)
	}

	resp := &pb.UserResponsesDTO{Message: "Get all users", Users: userResponses}
	return resp, nil
}
