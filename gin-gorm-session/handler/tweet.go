package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tanimutomo/go-samples/gin-gorm-session/db"
)

// Get a list of tweets
func GetTweets() gin.HandlerFunc {
	return func(c *gin.Context) {
		tweets := db.GetTweets()
		c.JSON(http.StatusOK, gin.H{"tweets": tweets})
	}
}

// Post a new tweet
func PostNewTweet() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tweet db.Tweet
		// Validation
		if err := c.Bind(&tweet); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err, "tweet": tweet})
		} else {
			db.PostTweet(tweet.Content)
			c.JSON(http.StatusOK, gin.H{"message": "Success to post a new tweet"})
		}
	}
}

// Check the detail of a tweet
func GetTweetDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		tweet := db.GetTweet(id)
		c.JSON(http.StatusOK, gin.H{"tweet": tweet})
	}
}

// Update
func UpdateTweet() gin.HandlerFunc {
	return func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		var tweet db.Tweet
		c.Bind(&tweet)
		db.UpdateTweet(id, tweet.Content)
		c.JSON(http.StatusOK, gin.H{"message": "Sccess to update a tweet"})
	}
}

// Delete
func DeleteTweet() gin.HandlerFunc {
	return func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		db.DeleteTweet(id)
		c.JSON(http.StatusFound, gin.H{"message": "Success to delete a tweet"})
	}
}
