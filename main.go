package main

import (
	"fmt"

	"dumbmerch/database"
	"dumbmerch/pkg/mysql"
	"dumbmerch/routes"
	"github.com/joho/godotenv" 

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	errEnv := godotenv.Load()
	if errEnv != nil{
		panic("failed to load env file")
	}

	mysql.DatabaseConnection()
	database.RunMigration()

	routes.RouteInit(e.Group("/api/v1"))
	e.Static("/uploads", "./uploads")
	
	fmt.Println("Running on port 5000")
	e.Logger.Fatal(e.Start("localhost:5000"))
}
