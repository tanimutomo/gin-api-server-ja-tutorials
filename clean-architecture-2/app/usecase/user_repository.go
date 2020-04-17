package usecase

import (
	"github.com/tanimutomo/go-samples/clean-architecture-2/app/domain"
)

type UserRepository interface {
	FindById(id int) (event domain.Users, err error)
}
