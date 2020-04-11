package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanimutomo/go-samples/gin-gorm-session/crypto"
	"github.com/tanimutomo/go-samples/gin-gorm-session/db"
)

// Signup
func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user db.User
		// Validation
		if err := c.Bind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		} else {
			// Check same username exists
			if err := db.CreateUser(user.Username, user.Password); len(err) != 0 {
				c.JSON(http.StatusBadRequest, gin.H{"Error": err})
			} else {
				c.JSON(http.StatusFound, gin.H{"message": "Success to signup"})
			}
		}
	}
}

// Login
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user db.User
		if err := c.Bind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		} else {
			// Get hashed password
			dbPassword := db.GetUser(user.Username).Password
			log.Println(dbPassword)
			sentPassword := user.Password
			// Compare user sent password to db password
			if err := crypto.CompareHashAndPassword(dbPassword, sentPassword); err != nil {
				log.Println("Failed to login")
				c.JSON(http.StatusBadRequest, gin.H{"Error": err})
			} else {
				log.Println("Success to login")
				c.JSON(http.StatusFound, gin.H{"message": "Success to login"})
			}
		}
	}
}
