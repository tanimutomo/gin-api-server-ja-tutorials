package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/tanimutomo/go-samples/clean-architecture-2/app/interfaces/controllers"
)

type Router struct {
	DB  *DB
	Gin *gin.Engine
}

func NewRouter(db *DB) *Router {
	r := &Router{
		DB:  db,
		Gin: gin.Default(),
	}
	r.setRouter()
	return r
}

func (r *Router) setRouter() {
	userController := controllers.NewUserController(r.DB)
	r.Gin.GET("/users/:id", func(c *gin.Context) { userController.Get(c) })
}

func (r *Router) Run(port string) {
	r.Gin.Run(port)
}
