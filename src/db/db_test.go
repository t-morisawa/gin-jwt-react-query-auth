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
	SignUp("morisawa"+random_str+"@exmaple.com", "password", "morisawa")

	// Read
	var user User
	db.Last(&user, "username = ?", "morisawa")

	// パスワードの一致確認
	err = compareHashAndPassword((&user).Password, "password")
	if err != nil {
		return 1, err
	}

	// Update
	db.Model(&user).Update("username", "morisawa1")
	db.Model(&user).Updates(User{Username: "morisawa2"})
	db.Model(&user).Updates(map[string]interface{}{"username": "morisawa3"})

	// Delete
	// gorm.Modelを使う場合は論理削除となる
	// https://gorm.io/ja_JP/docs/delete.html#%E8%AB%96%E7%90%86%E5%89%8A%E9%99%A4
	db.Where("username = ?", "morisawa3").Delete(&User{}) // find user with code D42

	return 0, nil
}

func TestDbConnect(t *testing.T) {
	result, err := dbConnect()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if result != 0 {
		t.Fatal("failed test")
	}
}
