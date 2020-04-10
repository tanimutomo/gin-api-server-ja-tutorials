package main

import (
	// "log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Declare tweet model
type Tweet struct {
	gorm.Model
	Content string `json:"content" binding:"required"`
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "test"
	PASS := "12345678"
	DBNAME := "test"
	// postfix 'parse...' for charcode of mysql
	CONNECT := USER + ":" + PASS + "@/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	return db
}

// Initialize DB
func dbInit() {
	db := gormConnect()

	defer db.Close()
	db.AutoMigrate(&Tweet{})
}

// Insert data
func dbInsert(content string) {
	db := gormConnect()

	defer db.Close()
	// Insert
	db.Create(&Tweet{Content: content})
}

// Update DB
func dbUpdate(id int, tweetText string) {
	db := gormConnect()
	var tweet Tweet
	db.First(&tweet, id)
	tweet.Content = tweetText
	db.Save(&tweet)
	db.Close()
}

// Get all data
func dbGetAll() []Tweet {
	db := gormConnect()

	defer db.Close()
	var tweets []Tweet
	// Get all tweet data by specifying empty condition as the Find argument
	db.Order("created_at desc").Find(&tweets)
	return tweets
}

// Get one
func dbGetOne(id int) Tweet {
	db := gormConnect()
	var tweet Tweet
	db.First(&tweet, id) // Get tweet and push to tweet's value
	db.Close()
	return tweet
}

// Delete a tweet
func dbDelete(id int) {
	db := gormConnect()
	var tweet Tweet
	db.First(&tweet, id)
	db.Delete(&tweet)
	db.Close()
}

func main() {
	r := gin.Default()

	dbInit()

	// Get a list of tweets
	r.GET("/tweets", func(c *gin.Context) {
		tweets := dbGetAll()
		c.JSON(http.StatusOK, gin.H{"tweets": tweets})
	})

	// Post a new tweet
	r.POST("/tweets", func(c *gin.Context) {
		var tweet Tweet
		// Validation
		if err := c.Bind(&tweet); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err, "tweet": tweet})
		} else {
			dbInsert(tweet.Content)
			c.JSON(http.StatusOK, gin.H{"message": "Success to post a new tweet"})
		}
	})

	// Check the detail of a tweet
	r.GET("/tweets/:id/detail", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		tweet := dbGetOne(id)
		c.JSON(http.StatusOK, gin.H{"tweet": tweet})
	})

	// Update
	r.POST("/tweets/:id/update", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		var tweet Tweet
		c.Bind(&tweet)
		dbUpdate(id, tweet.Content)
		c.JSON(http.StatusOK, gin.H{"message": "Sccess to update a tweet"})
	})

	// Delete
	r.POST("/tweets/:id/delete", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		dbDelete(id)
		c.JSON(http.StatusFound, gin.H{"message": "Success to delete a tweet"})
	})

	r.Run(":8080")
}
