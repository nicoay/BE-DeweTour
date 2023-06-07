package database

import (
	"dumbmerch/models"
	"dumbmerch/pkg/mysql"
	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.User{}, &models.Profile{}, &models.Country{}, &models.Tour{}, &models.Transaction{})
	if err != nil {
		panic("Migration Failed")
	}
	fmt.Println("Success Migration")
}
