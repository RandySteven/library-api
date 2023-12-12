package handler_grpc

import (
	"context"
	"log"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/pb"
)

type BorrowHandler struct {
	pb.UnimplementedBorrowBookServiceServer
	usecase interfaces.BorrowUseCase
}

func (h *BorrowHandler) CreateBorrowBook(ctx context.Context, req *pb.BorrowRequest) (*pb.BorrowBookResponse, error) {

	userId := ctx.Value("id").(uint)
	log.Println(userId)
	borrow := &models.Borrow{
		UserID: userId,
		BookID: uint(req.BookId),
	}

	borrow, err := h.usecase.CreateBorrowRecord(ctx, borrow)

	if err != nil {
		return nil, err
	}

	return &pb.BorrowBookResponse{
		Message: "Borrow success",
		BorrowResponse: &pb.BorrowResponse{
			Id:           uint32(borrow.ID),
			UserId:       uint32(borrow.UserID),
			BookId:       uint32(borrow.BookID),
			BorrowStatus: borrow.BorrowStatus,
			BookQuantity: uint32(borrow.Book.Quantity),
		},
	}, nil
}

func (h *BorrowHandler) RetrunBook(ctx context.Context, req *pb.ReturnRequest) (*pb.BorrowBookResponse, error) {

	userId := ctx.Value("id").(uint)
	borrow, err := h.usecase.ReturnBorrowedBookByBorrowId(ctx, uint(req.BorrowId), userId)
	if err != nil {
		return nil, err
	}

	return &pb.BorrowBookResponse{
		Message: "Return successs",
		BorrowResponse: &pb.BorrowResponse{
			Id:     uint32(borrow.ID),
			BookId: uint32(borrow.BookID),
			UserId: uint32(borrow.UserID),
		},
	}, nil
}

func NewBorrowHandler(usecase interfaces.BorrowUseCase) *BorrowHandler {
	return &BorrowHandler{usecase: usecase}
}
