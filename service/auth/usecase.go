package auth

import (
	"exchequer/service/auth/delivery/http/request"
	"exchequer/service/auth/delivery/http/response"
)

type Usecase interface {
	SignUp(request request.SignUpRequest) error
	SignIn(request request.SignInRequest) (response.AuthResponse, error)
}
