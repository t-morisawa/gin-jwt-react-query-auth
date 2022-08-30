package db

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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
	password, err := passwordEncrypt("password")
	if err != nil {
		return 1, errors.New("failed to encrypt password")
	}
	db.Create(&User{FirstName: "toma", LastName: "morisawa", Password: password, Email: "morisawa" + random_str + "@exmaple.com"})

	// Read
	var user User
	// db.First(&user, 1)                           // find user with integer primary key
	db.First(&user, "last_name = ?", "morisawa") // find user with code D42

	err = compareHashAndPassword((&user).Password, "password")
	if err != nil {
		return 1, errors.New("wrong password")
	}

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

func TestDbConnect(t *testing.T) {
	result, err := dbConnect()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if result != 1 {
		t.Fatal("failed test")
	}
}
