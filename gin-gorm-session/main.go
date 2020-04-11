package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tanimutomo/go-samples/gin-gorm-session/db"
	"github.com/tanimutomo/go-samples/gin-gorm-session/handler"
)

func main() {
	r := gin.Default()

	db.Init()

	// Signup
	r.POST("/signup", handler.Signup())
	// Login
	r.POST("/login", handler.Login())

	// Get a list of tweets
	r.GET("/tweets", handler.GetTweets())
	// Post a new tweet
	r.POST("/tweets", handler.PostNewTweet())
	// Check the detail of a tweet
	r.GET("/tweets/:id/detail", handler.GetTweetDetail())
	// Update
	r.POST("/tweets/:id/update", handler.UpdateTweet())
	// Delete
	r.POST("/tweets/:id/delete", handler.DeleteTweet())

	r.Run(":8080")
}
