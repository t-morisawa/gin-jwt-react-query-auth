package db

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	// タグは下記を参照
	// https://gorm.io/ja_JP/docs/models.html#%E3%83%95%E3%82%A3%E3%83%BC%E3%83%AB%E3%83%89%E3%81%AB%E6%8C%87%E5%AE%9A%E5%8F%AF%E8%83%BD%E3%81%AA%E3%82%BF%E3%82%B0
	gorm.Model
	Email    string `gorm:"size:255;unique"`
	Password string `gorm:"size:255"` // ハッシュ化されたパスワードを格納すること
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
		return err
	}
	result := db.Create(user)
	if result.Error != nil {
		return err
	}

	return nil
}

func SignUp(email string, passwordRaw string, username string) error {
	password, err := passwordEncrypt(passwordRaw)
	if err != nil {
		return err
	}
	err = CreateUser(&User{Username: username, Password: password, Email: email})
	if err != nil {
		return err
	}
	return nil
}

func GetUser(email string) (*User, error) {
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	var user User
	result := db.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, err
	}
	return &user, nil
}

func Login(email string, passwordRaw string) (*User, error) {
	user, err := GetUser(email)
	if err != nil {
		return nil, err
	}
	err = compareHashAndPassword(user.Password, passwordRaw)
	if err != nil {
		return nil, err
	}
	return user, nil
}
