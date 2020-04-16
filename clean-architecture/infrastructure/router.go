package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/tanimutomo/go-samples/clean-architecture/src/app/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	userController := controllers.NewUserController(NewSqlHander())

	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.GET("/users", func(c *gin.Context) { userController.Index(c) })
	router.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })

	Router = router
}
