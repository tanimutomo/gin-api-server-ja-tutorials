package usecase

import "github.com/tanimutomo/go-samples/clean-architecture/src/app/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(user domain.User) (err error) {
	_, err = interactor.UserRepository.Store(user)
	return
}

func (interactor *UserInteractor) Users() (users domain.Users, err error) {
	users, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) UserById(identifier int) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindById(identifier)
	return
}
