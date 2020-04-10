package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.POST("/login", func(c *gin.Context) {
		var user User
		c.Bind(&user)
		// Create a new session
		session := sessions.Default(c)
		session.Set(user.Username, user.Password)
		session.Save()
		c.JSON(http.StatusOK, gin.H{"message": "Success to login"})
		log.Println(user.Username, ":", session.Get(user.Username))
		log.Println("NotExistUser:", session.Get("as;dlkfjakl;sdj"))
	})
	r.POST("/logout", func(c *gin.Context) {
		var user User
		c.Bind(&user)
		// Clear session
		session := sessions.Default(c)
		session.Delete(user.Username)
		session.Save()
		c.JSON(http.StatusOK, gin.H{"message": "Success to logout"})
		log.Println(user.Username, ":", session.Get(user.Username))
	})
	r.Run(":8080")
}
