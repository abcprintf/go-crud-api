package db

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB
var err error

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key"`
	Username string    `gorm:"type:varchar(100);unique_index"`
	Password string    `gorm:"type:varchar(100)"`
}

func InitPostgresDB() {
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		dbUser   = os.Getenv("DB_USER")
		dbName   = os.Getenv("DB_NAME")
		password = os.Getenv("DB_PASSWORD")
	)
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, dbUser, dbName, password)
	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&User{})
}

func CreateUser(user *User) (*User, error) {
	user.ID = uuid.New()
	res := db.Create(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func GetUser(id string) (*User, error) {
	var user User
	res := db.First(&user, "id = ?", id)
	if res.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("user of id %s not found", id))
	}
	return &user, nil
}

func GetUsers() ([]*User, error) {
	var users []*User
	res := db.Find(&users)
	if res.Error != nil {
		return nil, errors.New("no users found")
	}
	return users, nil
}

func UpdateUser(user *User) (*User, error) {
	var userToUpdate User
	result := db.Model(&userToUpdate).Where("id = ?", user.ID).Updates(user)
	if result.RowsAffected == 0 {
		return &userToUpdate, errors.New("user not updated")
	}
	return user, nil
}

func DeleteUser(id string) error {
	var deletedUser User
	result := db.Where("id = ?", id).Delete(&deletedUser)
	if result.RowsAffected == 0 {
		return errors.New("user not deleted")
	}
	return nil
}
