package db

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/tanimutomo/go-samples/gin-gorm-session/crypto"
)

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
func Init() {
	db := gormConnect()

	defer db.Close()
	db.AutoMigrate(&Tweet{})
	db.AutoMigrate(&User{})
}

// Tweet //

// Declare tweet model
type Tweet struct {
	gorm.Model
	Content string `json:"content" binding:"required"`
}

// Insert data
func PostTweet(content string) {
	db := gormConnect()

	defer db.Close()
	// Insert
	db.Create(&Tweet{Content: content})
}

// Update DB
func UpdateTweet(id int, tweetText string) {
	db := gormConnect()
	var tweet Tweet
	db.First(&tweet, id)
	tweet.Content = tweetText
	db.Save(&tweet)
	db.Close()
}

// Get all data
func GetTweets() []Tweet {
	db := gormConnect()

	defer db.Close()
	var tweets []Tweet
	// Get all tweet data by specifying empty condition as the Find argument
	db.Order("created_at desc").Find(&tweets)
	return tweets
}

// Get one
func GetTweet(id int) Tweet {
	db := gormConnect()
	var tweet Tweet
	db.First(&tweet, id) // Get tweet and push to tweet's value
	db.Close()
	return tweet
}

// Delete a tweet
func DeleteTweet(id int) {
	db := gormConnect()
	var tweet Tweet
	db.First(&tweet, id)
	db.Delete(&tweet)
	db.Close()
}

// User //

type User struct {
	gorm.Model
	Username string `json:"username" binding:"required" gorm:"unique;not null"`
	Password string `json:"password" binding:"required"`
}

// Register a new user
func CreateUser(username string, password string) []error {
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
func GetUser(username string) User {
	db := gormConnect()
	var user User
	db.First(&user, "username = ?", username)
	db.Close()
	return user
}
