package server

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/configs"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/usecase"
)

type (
	Handlers struct {
		BookHandler   interfaces.BookHandler
		UserHandler   interfaces.UserHandler
		BorrowHandler interfaces.BorrowHandler
	}
)

func NewHandlers(repo configs.Repository) (*Handlers, error) {
	bookUsecase := usecase.NewBookUseCase(repo.BookRepository, repo.AuthorRepository)
	userUsecase := usecase.NewUserUseCase(repo.UserRepository)
	borrowUsecase := usecase.NewBorrowUseCase(repo.BorrowRepository)

	return &Handlers{
		BookHandler:   handler.NewBookHandler(bookUsecase),
		UserHandler:   handler.NewUserHandler(userUsecase),
		BorrowHandler: handler.NewBorrowHandler(borrowUsecase),
	}, nil
}

// func (h Handlers) RequestMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
// 		xRequestId := uuid.New().String()
// 		xTimestamp := time.Now()
// 		rctx := context.WithValue(req.Context(), enums.XRequestID, xRequestId)
// 		rctx = context.WithValue(rctx, enums.XRequestID, xRequestId)
// 		rctx2 := context.WithValue(rctx, enums.XTimestamp, xTimestamp)

// 		log.Printf("[Request] %v %v %s \n", req.Method, req.URL.Path, xRequestId)
// 		log.Printf("[Time] %v ", xTimestamp)

// 		next.ServeHTTP(res, req.WithContext(rctx2))
// 	})
// }

// func (h Handlers) LogMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
// 		log.Println("Log")
// 		next.ServeHTTP(res, req)
// 	})
// }

// func (h Handlers) RoleMiddleware(c *gin.Context) {
// 	claims := h.validateToken(c)
// 	if claims == nil {
// 		resp := response.Response{
// 			Errors: []string{"Unauthorized. Invalid token"},
// 		}
// 		utils.ResponseHandler(c.Writer, http.StatusUnauthorized, resp)
// 		c.Abort()
// 		return
// 	}

// 	if claims.User != nil && claims.User.RoleID != 1 {
// 		resp := response.Response{
// 			Errors: []string{"Access denied"},
// 		}
// 		utils.ResponseHandler(c.Writer, http.StatusForbidden, resp)
// 		c.Abort()
// 		return
// 	}

// 	c.Next()

// }

// func (h Handlers) validateToken(c *gin.Context) *configs.JWTClaim {
// 	tokenString, err := c.Cookie("token")
// 	if err != nil {
// 		return nil
// 	}

// 	claims := &configs.JWTClaim{}
// 	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
// 		return configs.JWT_KEY, nil
// 	})

// 	if err != nil || !token.Valid {
// 		return nil
// 	}

// 	return claims
// }

// func (h Handlers) AuthMiddleware(c *gin.Context) {
// 	claims := h.validateToken(c)
// 	if claims == nil {
// 		resp := response.Response{
// 			Errors: []string{"Unauthorized. Invalid token"},
// 		}
// 		utils.ResponseHandler(c.Writer, http.StatusUnauthorized, resp)
// 		c.Abort()
// 		return
// 	}

// 	c.Set("user", claims.User)
// 	c.Next()
// }
