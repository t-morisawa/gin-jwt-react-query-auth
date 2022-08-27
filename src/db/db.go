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
	user := User{FirstName: "toma", LastName: "morisawa", Email: "morisawa@exmaple.com", Status: true}
	db.Create(&user)

	// Read
	var product User
	// db.First(&product, 1)                           // find product with integer primary key
	db.First(&product, "last_name = ?", "morisawa") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("first_name", "toma")
	// Update - update multiple fields
	db.Model(&product).Updates(User{FirstName: "toma2", LastName: "morisawa2", Email: "morisawa@exmaple.com"})
	db.Model(&product).Updates(map[string]interface{}{"first_name": "toma3", "last_name": "morisawa3"})

	// Delete - delete product
	db.Delete(&product, 1)

	return 0, nil
}
