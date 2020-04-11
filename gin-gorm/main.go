package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/tanimutomo/go-samples/gin-gorm/crypto"
)

// Declare tweet model
type Tweet struct {
	gorm.Model
	Content string `json:"content" binding:"required"`
}

type User struct {
	gorm.Model
	Username string `json:"username" binding:"required" gorm:"unique;not null"`
	Password string `json:"password" binding:"required"`
}

func gormConnect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	DBMS := os.Getenv("mytweet_DBMS")
	USER := os.Getenv("mytweet_USER")
	PASS := os.Getenv("mytweet_PASS")
	DBNAME := os.Getenv("mytweet_DBNAME")
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
	db.AutoMigrate(&User{})
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

// Register a new user
func createUser(username string, password string) []error {
	passwordEncrypt, _ := crypto.PasswordEncrypt(password)
	db := gormConnect()
	defer db.Close()
	// Insert a new user to db
	if err := db.Create(
		&User{Username: username, Password: passwordEncrypt},
	).GetErrors(); err != nil {
		return err
	}
	return nil
}

// Find a user
func getUser(username string) User {
	db := gormConnect()
	var user User
	db.First(&user, "username = ?", username)
	db.Close()
	return user
}

func main() {
	r := gin.Default()

	dbInit()

	// Signup
	r.POST("/signup", func(c *gin.Context) {
		var user User
		// Validation
		if err := c.Bind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		} else {
			// Check same username exists
			if err := createUser(user.Username, user.Password); len(err) != 0 {
				c.JSON(http.StatusBadRequest, gin.H{"Error": err})
			} else {
				c.JSON(http.StatusFound, gin.H{"message": "Success to signup"})
			}
		}
	})

	// Login
	r.POST("/login", func(c *gin.Context) {
		var user User
		if err := c.Bind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		} else {
			// Get hashed password
			dbPassword := getUser(user.Username).Password
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
	})

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
