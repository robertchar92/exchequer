package usecase

import (
	"exchequer/models"
	"exchequer/service/book/delivery/http/request"

	"exchequer/utils/errors"

	"github.com/google/uuid"
)

func (u *Usecase) Create(userID string, username string, request request.BookCreateRequest) (*models.Book, error) {
	bookM := &models.Book{
		ID:       uuid.New().String(),
		UserID:   userID,
		Username: username,
		Name:     *request.Name,
		Balance:  0,
	}

	err := u.bookRepo.Insert(bookM, nil)
	if err != nil {
		err := errors.ErrUnprocessableEntity
		err.Message = "Failed to create book!"
		return nil, err
	}

	return bookM, nil
}
