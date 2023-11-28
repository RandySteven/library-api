package handler_test

// type AuthHandlerTestSuite struct {
// 	suite.Suite
// 	authUseCase *mocks.AuthUseCase
// 	authHandler *handler.AuthHandler
// 	router      *gin.Engine
// }

// func (suite *AuthHandlerTestSuite) Setup() {
// 	suite.authUseCase = mocks.NewAuthUseCase(suite.T())
// 	suite.authHandler = handler.NewAuthHandler(suite.authUseCase)
// 	suite.router = gin.Default()
// }

// func TestAuthHandler(t *testing.T) {
// 	suite.Run(t, new(AuthHandlerTestSuite))
// }

// func (suite *AuthHandlerTestSuite) TestLoginAuth() {
// 	suite.Run("should return 200 when success to login", func() {
// 		requestBody := `{
// 			"email": "randysteven12@gmail.com",
// 			"password": "test_1234"
// 		}`
// 		suite.authUseCase.On("LoginUserByEmail", mock.AnythingOfType("context.Context"), mock.Anything, mock.Anything).
// 			Return("token", nil)

// 		req, _ := http.NewRequest("POST", "/login", strings.NewReader(requestBody))
// 		w := httptest.NewRecorder()

// 		c, _ := gin.CreateTestContext(w)
// 		c.Request = req

// 		suite.router.POST("/login", suite.authHandler.LoginUser)
// 		suite.Equal(http.StatusOK, w.Code)
// 	})
// }

// func (suite *AuthHandlerTestSuite) TestRegisterAuth() {
// 	suite.Run("should return 201 when user succes to register", func() {
// 		registerBody := `{
// 			"name": "Randy Steven",
// 			"email": "randysteven12@gmail.com",
// 			"password": "test_1234",
// 			"phone_number": "0123456789"
// 		}`
// 		suite.authUseCase.On("RegisterUser", mock.AnythingOfType("context.Context"), mock.AnythingOfType("*models.User")).
// 			Return(mock.AnythingOfType("*models.User"), nil)
// 		req, _ := http.NewRequest("POST", "/register", strings.NewReader(registerBody))
// 		w := httptest.NewRecorder()

// 		c, _ := gin.CreateTestContext(w)
// 		c.Request = req

// 		suite.router.POST("/register", suite.authHandler.RegisterUser)
// 		suite.Equal(http.StatusCreated, w.Code)
// 	})
// }
