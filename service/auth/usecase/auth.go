package usecase

import (
	"exchequer/service/auth"
	"exchequer/vendor/github.com/golang-jwt/jwt/v4"
)

type Usecase struct {
}

type authClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Role     string `json:"role"`
	Id       uint64 `json:"id"`
	IssuedAt int64  `json:"orig_iat,omitempty"`
}

func New() auth.Usecase {
	return &Usecase{}
}
