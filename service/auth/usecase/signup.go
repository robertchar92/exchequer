package usecase

import (
	"exchequer/models"
	"exchequer/service/auth/delivery/http/request"
	"exchequer/utils/errors"
	"exchequer/utils/helpers"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (u *Usecase) SignUp(request request.SignUpRequest) error {
	// to lower cap email and username
	request.Username = strings.ReplaceAll(strings.ToLower(request.Username), " ", "")
	isAlpha := regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	if !isAlpha(request.Username) {
		err := errors.ErrUnprocessableEntity
		err.Message = "username must be a character, number or underscore"

		return err
	}

	request.Email = strings.ReplaceAll(strings.ToLower(request.Email), " ", "")
	if !helpers.CheckIsEmail(request.Email) {
		err := errors.ErrUnprocessableEntity
		err.Message = "email is not a valid email"

		return err
	}

	hasSpecialChar := regexp.MustCompile(`[!@#$%^&*()\-_=+{}\[\]:";'<>,.?/]`).MatchString(request.Password)
	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(request.Password)
	hasLowercase := regexp.MustCompile(`[a-z]`).MatchString(request.Password)
	hasNumber := regexp.MustCompile(`\d`).MatchString(request.Password)

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
		err.Message = fmt.Sprint("Password must contain ", s)

		return err
	}

	bannedUsernames := []string{"admin"}

	for _, bannedUsername := range bannedUsernames {
		if strings.Contains(request.Username, bannedUsername) {
			err := errors.ErrUnprocessableEntity
			err.Message = fmt.Sprint("username may not contain any of :", strings.Join(bannedUsernames, ","))
			return err
		}
	}

	userM := &models.User{
		Username: request.Username,
		Name:     request.Name,
		Email:    request.Email,
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error-encrypting-password:", err)
		return errors.ErrUnprocessableEntity
	}
	userM.Password = string(password)

	userM.ID = uuid.New().String()

	err = u.userRepo.Insert(userM, nil)
	if err != nil {
		return err
	}

	return nil
}
