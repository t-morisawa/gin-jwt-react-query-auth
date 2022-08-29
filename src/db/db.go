package db

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"size:255"`
	LastName  string `gorm:"size:255"`
	Email     string `gorm:"size:255"`
	Status    bool
}

const DSN = "user:password@tcp(gin-jwt-react-query-auth_db_1:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"

func CreateUser(user *User) error {
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		return errors.New("failed to connect database")
	}
	db.Create(user)

	return nil
}

func dbConnect() (int, error) {
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		return 1, errors.New("failed to connect database")
	}

	// マイグレーション
	// 内容によってはsqlファイルの初期作成を上書きする
	db.AutoMigrate(&User{})

	// ユーザ作成
	rand.Seed(time.Now().UnixNano())
	random_str := fmt.Sprint(rand.Intn(9999))
	db.Create(&User{FirstName: "toma", LastName: "morisawa", Email: "morisawa" + random_str + "@exmaple.com", Status: true})

	// Read
	var user User
	// db.First(&user, 1)                           // find user with integer primary key
	db.First(&user, "last_name = ?", "morisawa") // find user with code D42

	// Update - update user's price to 200
	db.Model(&user).Update("first_name", "toma")

	// Update - update multiple fields
	db.Model(&user).Updates(User{FirstName: "toma2"})
	db.Model(&user).Updates(map[string]interface{}{"first_name": "toma3"})

	// gorm.Modelを使う場合は論理削除となる
	// https://gorm.io/ja_JP/docs/delete.html#%E8%AB%96%E7%90%86%E5%89%8A%E9%99%A4
	db.Where("last_name = ?", "morisawa").Delete(&User{}) // find user with code D42

	return 0, nil
}
