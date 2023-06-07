package routes

import (
	"dumbmerch/handlers"

	"github.com/labstack/echo/v4"
)

func ProductRoute(e *echo.Group) {
	e.GET("/products", handlers.FindProduct)
}
