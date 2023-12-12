package handler_grpc

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/pb"
)

type BookHandler struct {
	usecase interfaces.BookUseCase
	pb.UnimplementedBookServiceServer
}

func NewBookHandler(usecase interfaces.BookUseCase) *BookHandler {
	return &BookHandler{usecase: usecase}
}

func (h *BookHandler) GetAllBooks(ctx context.Context, empty *pb.Empty) (*pb.GetBookResponses, error) {

	// id := ctx.Value("id")

	// if id == "" {
	// 	return nil, &apperror.ErrUnauthorized{}
	// }

	books, err := h.usecase.GetAllBooks(ctx, nil)
	if err != nil {
		return nil, err
	}

	bookResponses := []*pb.BookResponse{}
	for _, book := range books {
		bookRes := pb.BookResponse{
			Id:       uint32(book.ID),
			Title:    book.Title,
			Quantity: uint32(book.Quantity),
			Cover:    book.Cover,
			Author: &pb.Author{
				Id:        uint32(book.AuthorID),
				Name:      book.Author.Name,
				CreatedAt: book.CreatedAt.GoString(),
				UpdatedAt: book.UpdatedAt.GoString(),
				DeletedAt: book.DeletedAt.Time.GoString(),
			},
			CreatedAt: book.CreatedAt.GoString(),
			UpdatedAt: book.UpdatedAt.GoString(),
			DeletedAt: book.DeletedAt.Time.String(),
		}

		bookResponses = append(bookResponses, &bookRes)
	}

	return &pb.GetBookResponses{
		Message:       "Get all books",
		BookResponses: bookResponses,
	}, nil
}

func (h *BookHandler) CreateBook(ctx context.Context, req *pb.BookRequest) (*pb.BookResponse, error) {
	book := &models.Book{
		Title:       req.Title,
		Quantity:    uint(req.Quantity),
		Description: req.Description,
		Cover:       req.Cover,
		AuthorID:    uint(req.AuthorId),
	}

	book, err := h.usecase.CreateBook(ctx, book)
	if err != nil {
		return nil, err
	}

	return &pb.BookResponse{
		Id:          uint32(book.ID),
		Title:       book.Title,
		Description: book.Description,
		Quantity:    uint32(book.Quantity),
		Cover:       book.Cover,
		Author:      nil,
		CreatedAt:   book.CreatedAt.GoString(),
		UpdatedAt:   book.UpdatedAt.GoString(),
	}, nil
}
