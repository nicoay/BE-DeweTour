package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repository"

	"github.com/labstack/echo/v4"
)

func UsersRoute(e *echo.Group) {
	userRepository := repository.RepositoryUser(mysql.DB)
	h := handlers.HandleUser(userRepository)
	e.GET("/users", h.FindUsers)
	e.POST("/user", h.CreateUser)
	e.GET("/user/:id", h.GetUserById)
	e.PATCH("/user/:id", h.UpdateUser)
	e.DELETE("/user/:id", h.DeleteUser)
}

// func ProductRoute(e *echo.Group){
// 	e.GET("/products",handlers.FindProduct)
// }
