package server

import (
	"net/http"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/middleware"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) InitRouter(r *gin.RouterGroup) {

	// r.Use(func(ctx *gin.Context) {
	// 	go func() {
	// 		<-ctx.Done()
	// 		err := ctx.Request.Context().Err()
	// 		log.Println("contex done ", err)
	// 		ctx.Error(err)
	// 		if errors.Is(err, context.Canceled) {
	// 			log.Println("context cancelled")
	// 		}

	// 		if errors.Is(err, context.DeadlineExceeded) {
	// 			log.Println("context deadline exceeds")
	// 		}

	// 	}()
	// 	ctx.Next()
	// })

	r.Use(middleware.WithTimeOut)

	r.GET("/hello", func(ctx *gin.Context) {
		time.Sleep(time.Second * 5)
		ctx.JSON(http.StatusOK, gin.H{"hello": "world"})
	})

	bookRouter := r.Group("/books")
	bookRouter.Use(middleware.AuthMiddleware)

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

	borrowRouter := r.Group("/borrowing-records")
	borrowRouter.Use(middleware.AuthMiddleware)
	borrowRouter.POST("", h.BorrowHandler.CreateBorrowRecord)
	borrowRouter.GET("", h.BorrowHandler.GetAllBorrowsRecord)
	borrowRouter.PUT("/:id", h.BorrowHandler.ReturnBorrowBook)
}
