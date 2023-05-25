package usecase

import (
	"exchequer/models"
	"exchequer/service/auth"
	"exchequer/service/user"
	"exchequer/utils/errors"
	"log"
	"os"
	"time"

	"exchequer/lib/database_transaction"

	"github.com/golang-jwt/jwt/v4"
)

type Usecase struct {
	userRepo           user.Repository
	transactionManager database_transaction.Client
}

type authClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Role     string `json:"role"`
	Id       string `json:"id"`
	IssuedAt int64  `json:"orig_iat,omitempty"`
}

func New(
	userRepo user.Repository,
	transactionManager database_transaction.Client,
) auth.Usecase {
	return &Usecase{
		userRepo:           userRepo,
		transactionManager: transactionManager,
	}
}

func generateUserTokenString(userM *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, authClaims{
		Id:       userM.ID,
		Username: userM.Username,
		Role:     models.RoleNameUser,
		IssuedAt: time.Now().Unix(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1800000).Unix(),
		},
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Println("error-encoding-token", err)
		return "", errors.ErrUnprocessableEntity
	}

	return tokenString, nil
}
