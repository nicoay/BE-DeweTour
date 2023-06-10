package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repository"

	"github.com/labstack/echo/v4"
)

func UsersRoute(e *echo.Group) {
	userRepository := repository.RepositoryUser(mysql.DB)
	h := handlers.HandleUser(userRepository)
	e.GET("/users", middleware.Auth(h.FindUsers))
	e.POST("/user", h.CreateUser)
	e.GET("/user/:id", middleware.Auth(h.GetUserById))
	e.PATCH("/user/:id", middleware.Auth(middleware.UploadFile(h.UpdateUser)))
	e.DELETE("/user/:id", middleware.Auth(h.DeleteUser))
}

// func ProductRoute(e *echo.Group){
// 	e.GET("/products",handlers.FindProduct)
// }
