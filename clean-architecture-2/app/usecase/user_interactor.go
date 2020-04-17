package usecase

import (
	"github.com/tanimutomo/go-samples/clean-architecture-2/app/domain"
)

type UserInteractor struct {
	User UserRepository
	StatusCode int
}

func (interactor *UserInteractor) Get(id int) (user domain.UserForGet, err error) {
	// Get User
	foundUser, err := interactor.User.FindByID(id)
	if err != nil {
		interactor.StatusCode = 404
		return domain.UserForGet{}, err
	}
	user = foundUser.BuildForAGet()
	interactor.StatusCode = 200
	return user, nil
}