package main

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

var secreteKey = "75c92a074c341e9964329c0550c2673730ed8479c885c43122c90a2843177d5ef21cb50cfadcccb20aeb730487c11e09ee4dbbb02387242ef264e74cbee97213"

func main() {
	r := gin.Default()

	r.GET("/api/", func(c *gin.Context) {
		// Specify algorithm
		token := jwt.New(jwt.GetSigningMethod("HS256"))
		token.Claims = jwt.MapClaims{
			"user": "guest",
			"exp":  time.Now().Add(time.Hour * 1).Unix(),
		}

		// Give signature to token
		tokenString, err := token.SignedString([]byte(secreteKey))
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"token": tokenString})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token"})
		}
	})

	r.GET("/api/private/", func(c *gin.Context) {
		// Verrify signature
		token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			b := []byte(secreteKey)
			return b, nil
		})

		if err == nil {
			claims := token.Claims.(jwt.MapClaims)
			msg := fmt.Sprintf("Hello, '%s'", claims["user"])
			c.JSON(http.StatusOK, gin.H{"message": msg})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprint(err)})
		}
	})

	r.Run(":8080")
}
