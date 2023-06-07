package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repository"

	"github.com/labstack/echo/v4"
)

func CountriesRoute(e *echo.Group) {
	countryRepository := repository.RepositoryCountry(mysql.DB)
	h := handlers.HandleCountry(countryRepository)
	e.GET("/countries", h.FindCountries)
	e.POST("/country", h.CreateCountry)
	e.GET("/country/:id", h.GetCountry)
	e.PATCH("/country/:id", h.UpdateCountry)
	// e.DELETE("/user/:id", h.DeleteUser)
}

// func ProductRoute(e *echo.Group){
// 	e.GET("/products",handlers.FindProduct)
// }
