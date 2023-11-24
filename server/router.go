package server

import (
	"github.com/gin-gonic/gin"
)

func (h *Handlers) InitRouter(r *gin.RouterGroup) {

	bookRouter := r.Group("/books")
	bookRouter.POST("", h.BookHandler.CreateBook)
	bookRouter.GET("", h.BookHandler.GetAllBooks)
	// bookRouter.DELETE("/:id", h.BookHandler.DeleteProductById)
	// bookRouter.GET("/:id", h.BookHandler.GetProductById)
	// bookRouter.PUT("/:id", h.BookHandler.UpdateProductById)

	userRouter := r.Group("/users")
	userRouter.POST("", h.UserHandler.CreateUser)
	userRouter.GET("", h.UserHandler.GetAllUsers)
	// userRouter.Use(h.RoleMiddleware)
	// userRouter.HandleFunc("", h.UserHandler.GetAllUsers).Methods(http.MethodGet)
	// userRouter.HandleFunc("/{id}", h.UserHandler.GetUserById).Methods(http.MethodGet)

	borrowRouter := r.Group("/borrows")
	borrowRouter.POST("", h.BorrowHandler.CreateBorrowRecord)
	borrowRouter.GET("", h.BorrowHandler.GetAllBorrowsRecord)
	borrowRouter.PUT("/:id", h.BorrowHandler.ReturnBorrowBook)
}
