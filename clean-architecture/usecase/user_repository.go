package usecase

import "github.com/tanimutomo/gin-api-server-ja-tutorials/clean-architecture/src/app/domain"

type UserRepository interface {
	Store(domain.User) (int, error)
	FindById(int) (domain.User, error)
	FindAll() (domain.Users, error)
}
