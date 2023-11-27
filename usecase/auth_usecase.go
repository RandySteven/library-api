package usecase

import (
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/configs"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/entity/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/shared-projects/library-api/interfaces"
	"github.com/golang-jwt/jwt/v5"
)

type authUseCase struct {
	repo interfaces.AuthRepository
}

// LoginUserByEmail implements interfaces.authUseCase.
func (service *authUseCase) LoginUserByEmail(email string, password string) (string, error) {
	user, err := service.repo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	// log.Println(user.Password)
	// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	// if err != nil {
	// 	return "", err
	// }
	expTime := time.Now().Add(time.Minute * 60)
	claims := &configs.JWTClaim{
		UserId: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "issuer",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return tokenAlgo.SignedString(configs.JWT_KEY)
}

// RegisterUser implements interfaces.authUseCase.
func (service *authUseCase) RegisterUser(user *models.User) (*models.User, error) {
	return service.repo.RegisterUser(user)
}

func NewAuthUseCase(repo interfaces.AuthRepository) *authUseCase {
	return &authUseCase{repo: repo}
}

var _ interfaces.AuthUseCase = &authUseCase{}
