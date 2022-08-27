package db

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// id         int16
	FirstName string
	LastName  string
	Email     string
	Status    bool
}

func dbConnect() (int, error) {
	dsn := "user:password@tcp(gin-jwt-react-query-auth_db_1:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return 1, errors.New("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})

	// Create
	db.Create(&User{FirstName: "toma", LastName: "morisawa", Email: "morisawa@exmaple.com", Status: true})

	// Read
	var user User
	// db.First(&user, 1)                           // find user with integer primary key
	db.First(&user, "last_name = ?", "morisawa") // find user with code D42

	// Update - update user's price to 200
	db.Model(&user).Update("first_name", "toma")
	// Update - update multiple fields
	db.Model(&user).Updates(User{FirstName: "toma2", LastName: "morisawa2", Email: "morisawa2@exmaple.com"})
	db.Model(&user).Updates(map[string]interface{}{"first_name": "toma3", "last_name": "morisawa3", "email": "morisawa3@exmaple.com"})

	// Delete - delete user
	db.Delete(&user, "last_name = ?", "morisawa") // find user with code D42

	return 0, nil
}
