package usecase

import (
	"exchequer/models"
	"exchequer/utils/errors"
)

func (u *Usecase) Show(bookID string, userID string) (*models.Book, error) {
	bookM, err := u.bookRepo.FindByID(bookID)
	if err != nil {
		err := errors.ErrNotFound
		err.Message = "Failed to find book!"
		return nil, err
	}

	if userID != bookM.UserID {
		err := errors.ErrUnauthorized
		err.Message = "You are not the owner of the book!"
		return nil, err
	}

	return bookM, nil
}
