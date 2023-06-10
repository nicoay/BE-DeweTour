package mysql

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

//connecting database

func DatabaseConnection() {
	var err error
	var DB_HOST = "localhost"
	var DB_USER = "postgres"
	var DB_PASSWORD = "167669123"
	var DB_NAME = "dumbmerch"
	var DB_PORT = "5432"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Success to Connect Database")
}

// "{USER}:{PASSWORD}@tcp({HOST}:{PORT})/{DATABASE}?charset=utf8mb4&parseTime=True&loc=Local"
