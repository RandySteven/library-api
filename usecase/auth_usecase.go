package usecase

import (
	"context"
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
func (service *authUseCase) LoginUserByEmail(ctx context.Context, email string, password string) (string, error) {
	user, err := service.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	// log.Println(user.Password)
	// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	// if err != nil {
	// 	return "", err
	// }
	expTime := time.Now().Add(time.Minute * 15)
	claims := &configs.JWTClaim{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "issuer",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return tokenAlgo.SignedString(configs.JWT_KEY)
}

// RegisterUser implements interfaces.authUseCase.
func (service *authUseCase) RegisterUser(ctx context.Context, user *models.User) (*models.User, error) {
	return service.repo.RegisterUser(ctx, user)
}

func NewAuthUseCase(repo interfaces.AuthRepository) *authUseCase {
	return &authUseCase{repo: repo}
}

var _ interfaces.AuthUseCase = &authUseCase{}
