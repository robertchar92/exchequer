package usecase

import (
	"exchequer/models"
	"exchequer/service/user/delivery/http/request"
	"exchequer/utils/errors"
	"fmt"
	"regexp"

	"github.com/jinzhu/copier"

	"golang.org/x/crypto/bcrypt"
)

func (u *Usecase) Update(userID string, request request.UserUpdateRequest) (*models.User, error) {
	userM, _ := u.userRepo.FindByID(userID)

	_ = copier.Copy(userM, &request)

	if request.NewPassword != nil {
		if request.OldPassword == nil {
			err := errors.ErrUnprocessableEntity
			err.Message = "Please input old password to change current password"
			return nil, err
		}

		err := bcrypt.CompareHashAndPassword([]byte(userM.Password), []byte(*request.OldPassword))
		if err != nil {
			err := errors.ErrUnprocessableEntity
			err.Message = "Your old password is incorrect"
			return nil, err
		}

		hasSpecialChar := regexp.MustCompile(`[!@#$%^&*()\-_=+{}\[\]:";'<>,.?/]`).MatchString(*request.NewPassword)
		hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(*request.NewPassword)
		hasLowercase := regexp.MustCompile(`[a-z]`).MatchString(*request.NewPassword)
		hasNumber := regexp.MustCompile(`\d`).MatchString(*request.NewPassword)

		if !hasSpecialChar || !hasUppercase || !hasLowercase || !hasNumber {
			s := ""

			if !hasSpecialChar {
				s = s + "special characters"
			}
			if !hasUppercase {
				s = s + ", upper case characters"
			}
			if !hasLowercase {
				s = s + ", lower case characters"
			}
			if !hasNumber {
				s = s + ", number"
			}

			err := errors.ErrUnprocessableEntity
			err.Message = fmt.Sprint("New Password must contain ", s)

			return nil, err
		}

		newPassword, _ := bcrypt.GenerateFromPassword([]byte(*request.NewPassword), bcrypt.DefaultCost)
		userM.Password = string(newPassword)
	}

	err := u.userRepo.Update(userM, nil)
	if err != nil {
		return nil, err
	}

	return userM, nil
}
