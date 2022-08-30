package db

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"size:255"`
	Password  string `gorm:"size:255"`
	LastName  string `gorm:"size:255"`
	Email     string `gorm:"size:255"`
}

const DSN = "user:password@tcp(gin-jwt-react-query-auth_db_1:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"

func passwordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func compareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func CreateUser(user *User) error {
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		return errors.New("failed to connect database")
	}
	db.Create(user)

	return nil
}
