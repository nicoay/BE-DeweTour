package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	AuthRoutes(e)
	UsersRoute(e)
	ProductRoute(e)
	CountriesRoute(e)
	ToursRoute(e)
	TransRoute(e)
}