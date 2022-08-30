package db

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"size:255"`
	Password string `gorm:"size:255"`
	Username string `gorm:"size:255"`
}

const DSN = "user:password@tcp(gin-jwt-react-query-auth_db_1:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"

func passwordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func compareHashAndPassword(hash, passwordRaw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwordRaw))
}

func CreateUser(user *User) error {
	// パスワードはハッシュ化されている想定
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		return errors.New("failed to connect database")
	}
	db.Create(user)

	return nil
}

func GetUser(email string) (*User, error) {
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed to connect database")
	}
	var user User
	db.First(&user, "email = ?", email)

	return &user, nil
}

func Login(email string, passwordRaw string) (*User, error) {
	user, err := GetUser(email)
	if err != nil {
		return nil, errors.New("failed to login")
	}
	err = compareHashAndPassword(user.Password, passwordRaw)
	if err != nil {
		return nil, errors.New("wrong password")
	}
	return user, nil
}
