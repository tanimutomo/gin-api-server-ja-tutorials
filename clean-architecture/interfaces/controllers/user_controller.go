package controllers

import (
	"strconv"

	"github.com/tanimutomo/go-samples/clean-architecture/src/app/domain"
	"github.com/tanimutomo/go-samples/clean-architecture/src/app/interfaces/database"
	"github.com/tanimutomo/go-samples/clean-architecture/src/app/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				sqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c Context) {
	u := domain.User{}
	c.Bind(&u)
	err := controller.Interactor.Add(u)
	fi err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201)
}

func (controller *UserController) Index(c Context) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
}

func (controller *UserController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := contoroller.Interactor.UserById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
}
