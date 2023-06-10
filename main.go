package main

import (
	"fmt"

	"dumbmerch/database"
	"dumbmerch/pkg/mysql"
	"dumbmerch/routes"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.POST, echo.PATCH, echo.GET, echo.DELETE},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))

	godotenv.Load()

	mysql.DatabaseConnection()
	database.RunMigration()

	routes.RouteInit(e.Group("/api/v1"))
	e.Static("/uploads", "./uploads")

	fmt.Println("Running on port 5000")
	e.Logger.Fatal(e.Start("localhost:5000"))
}
