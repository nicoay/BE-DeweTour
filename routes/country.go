package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repository"

	"github.com/labstack/echo/v4"
)

func CountriesRoute(e *echo.Group) {
	countryRepository := repository.RepositoryCountry(mysql.DB)
	h := handlers.HandleCountry(countryRepository)
	e.GET("/countries", h.FindCountries)
	e.POST("/country", middleware.Auth(h.CreateCountry))
	e.GET("/country/:id", h.GetCountry)
	e.PATCH("/country/:id", middleware.Auth(h.UpdateCountry))
	e.DELETE("/country/:id", middleware.Auth(h.DeleteCountry))
}

// func ProductRoute(e *echo.Group){
// 	e.GET("/products",handlers.FindProduct)
// }
