package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

//connecting database

func DatabaseConnection() {
	var err error
	dsn := "root:@tcp(localhost:3306)/dumbmerch?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Success to Connect Database")
}

// "{USER}:{PASSWORD}@tcp({HOST}:{PORT})/{DATABASE}?charset=utf8mb4&parseTime=True&loc=Local"
