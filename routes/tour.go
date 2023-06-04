package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repository"

	"github.com/labstack/echo/v4"
)

func ToursRoute(e *echo.Group) {
	tourRepository := repository.RepositoryTour(mysql.DB)
	h := handlers.HandleTour(tourRepository)
	e.GET("/tours", h.FindTours)
	e.POST("/tour", middleware.Auth(middleware.UploadFile(h.CreateTour)))
	e.GET("/tour/:id", h.GetTour)
	e.PATCH("/tour/:id", middleware.Auth(h.UpdateTour))
	e.DELETE("/tour/:id", middleware.Auth(h.DeleteTour))
}
