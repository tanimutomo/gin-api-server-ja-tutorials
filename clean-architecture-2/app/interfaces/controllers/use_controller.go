package controllers

import (
	"strconv"

	"github.com/tanimutomo/go-samples/clean-architecture-2/app/interfaces/database"
	"github.com/tanimutomo/go-samples/clean-architecture-2/app/usecase"
)

type UserController struct {
	Interactor usecase.UserController
}

func NewUserController(db database.DB) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			User: &database.UserRepository{DB: db},
		},
	}
}

func (controller *UserController) Get(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.Get(id)
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewH{err.Error(), nil})
		return
	}
	c.JSON(controller.Interactor.StatusCode, NewH{"success", user})
}
