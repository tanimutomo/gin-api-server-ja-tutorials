package database

import (
	"errors"

	"github.com/tanimutomo/go-samples/clean-architecture-2/app/domain"
)

type UserRepository struct {
	DB DB
}

func (repo *UserRepository) FindByID(id int) (user domain.Users, err error) {
	user = domain.Users{}
	repo.DB.First(&user, id)
	if user.ID <= 0 {
		return domain.Users{}, errors.New("user is not found")
	}
	return user, nil
}
