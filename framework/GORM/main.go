package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID   int    `gorm:"primaryKey;column:id"`
	Name string `gorm:"size:255;column:name"`
	Age  int    `gorm:"column:age"`
}

func main() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open db: %v", err)
	}

	db.AutoMigrate(&User{})

	user := User{Name: "tutu", Age: 114514}
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatalf("Failed to create user: %v", result.Error)
	}

	var queriedUser User
	result = db.First(&queriedUser, user.ID)
	if result.Error != nil {
		log.Fatalf("Failed to query user: %v", result.Error)
	} else {
		fmt.Printf("Queried User: %+v\n", queriedUser)
	}

	queriedUser.Age = 20
	result = db.Save(&queriedUser)
	if result.Error != nil {
		log.Fatalf("Failed to update user: %v", result.Error)
	}

	// result = db.Delete(&queriedUser)
	// if result.Error != nil {
	// 	log.Fatalf("Failed to delete user: %v", result.Error)
	// }

}
