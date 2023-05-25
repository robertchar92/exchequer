package usecase

import (
	"exchequer/models"
	"exchequer/service/auth/delivery/http/request"
	"exchequer/service/auth/delivery/http/response"
	"exchequer/utils/errors"
	"exchequer/utils/helpers"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (u *Usecase) SignIn(request request.SignInRequest) (response.AuthResponse, error) {
	var userM *models.User

	request.UsernameOrEmail = strings.ToLower(request.UsernameOrEmail)
	if helpers.CheckIsEmail(request.UsernameOrEmail) {
		userM, _ = u.userRepo.FindByEmail(request.UsernameOrEmail)
		if userM == nil {
			err := errors.ErrFailedAuthentication
			err.Message = "Your username or password doesn't match!"
			return response.AuthResponse{}, err
		}
	} else {
		userM, _ = u.userRepo.FindByUsername(request.UsernameOrEmail)
		if userM == nil {
			err := errors.ErrFailedAuthentication
			err.Message = "Your username or password doesn't match!"
			return response.AuthResponse{}, err
		}
	}

	err := bcrypt.CompareHashAndPassword([]byte(userM.Password), []byte(request.Password))
	if err != nil {
		err := errors.ErrFailedAuthentication
		err.Message = "Your username or password doesn't match!"
		return response.AuthResponse{}, err
	}

	now := time.Now()

	userM.LastLoginAt = &now

	_ = u.userRepo.Update(userM, nil)

	// generate jwt token
	tokenString, err := generateUserTokenString(userM)
	if err != nil {
		return response.AuthResponse{}, err
	}

	return response.AuthResponse{
		Token: tokenString,
		User:  *userM,
	}, nil
}
