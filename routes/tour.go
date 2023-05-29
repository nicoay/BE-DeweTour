package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repository"

	"github.com/labstack/echo/v4"
)

func ToursRoute(e *echo.Group) {
	tourRepository := repository.RepositoryTour(mysql.DB)
	h := handlers.HandleTour(tourRepository)
	e.GET("/tours", h.FindTours)
	e.POST("/tour", h.CreateTour)
	e.GET("/tour/:id", h.GetTour)
	e.PATCH("/tour/:id", h.UpdateTour)
	// e.DELETE("/user/:id", h.DeleteUser)
}
