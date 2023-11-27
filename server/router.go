package server

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/middleware"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) InitRouter(r *gin.RouterGroup) {

	bookRouter := r.Group("/books")
	bookRouter.POST("", middleware.ErrorMiddleware(), h.BookHandler.CreateBook)
	bookRouter.GET("", middleware.ErrorMiddleware(), h.BookHandler.GetAllBooks)
	// bookRouter.DELETE("/:id", h.BookHandler.DeleteProductById)
	// bookRouter.GET("/:id", h.BookHandler.GetProductById)
	// bookRouter.PUT("/:id", h.BookHandler.UpdateProductById)

	userRouter := r.Group("/users")
	userRouter.POST("", middleware.ErrorMiddleware(), h.UserHandler.CreateUser)
	userRouter.GET("", middleware.ErrorMiddleware(), h.UserHandler.GetAllUsers)
	// userRouter.Use(h.RoleMiddleware)
	// userRouter.HandleFunc("", h.UserHandler.GetAllUsers).Methods(http.MethodGet)
	// userRouter.HandleFunc("/{id}", h.UserHandler.GetUserById).Methods(http.MethodGet)

	borrowRouter := r.Group("/borrowing-records")
	borrowRouter.POST("", middleware.ErrorMiddleware(), h.BorrowHandler.CreateBorrowRecord)
	borrowRouter.GET("", middleware.ErrorMiddleware(), h.BorrowHandler.GetAllBorrowsRecord)
	borrowRouter.PUT("/:id", middleware.ErrorMiddleware(), h.BorrowHandler.ReturnBorrowBook)
}
