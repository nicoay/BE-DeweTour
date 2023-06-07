package main

import (
	"fmt"

	"dumbmerch/database"
	"dumbmerch/pkg/mysql"
	"dumbmerch/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	mysql.DatabaseConnection()
	database.RunMigration()

	routes.RouteInit(e.Group("/api/v1"))
	fmt.Println("Running on port 5000")
	e.Logger.Fatal(e.Start("localhost:5000"))
}
