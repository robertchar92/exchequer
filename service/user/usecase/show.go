package usecase

import "exchequer/models"

func (u *Usecase) Show(userID string, unscoped ...bool) (*models.User, error) {
	userM, err := u.userRepo.FindByID(userID, unscoped...)
	if err != nil {
		return nil, err
	}

	return userM, nil
}
