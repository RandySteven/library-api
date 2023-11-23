package server

import (
	"github.com/gin-gonic/gin"
)

func (h *Handlers) InitRouter(r *gin.RouterGroup) {

	bookRouter := r.Group("/books")
	// bookRouter.POST("", h.BookHandler.CreateProduct)
	bookRouter.GET("", h.BookHandler.GetAllBooks)
	// bookRouter.DELETE("/:id", h.BookHandler.DeleteProductById)
	// bookRouter.GET("/:id", h.BookHandler.GetProductById)
	// bookRouter.PUT("/:id", h.BookHandler.UpdateProductById)

	// userRouter := r.PathPrefix("/users").Subrouter()
	// userRouter.Use(h.RoleMiddleware)
	// userRouter.HandleFunc("", h.UserHandler.GetAllUsers).Methods(http.MethodGet)
	// userRouter.HandleFunc("/{id}", h.UserHandler.GetUserById).Methods(http.MethodGet)
}
