package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repository"

	"github.com/labstack/echo/v4"
)

func TransRoute(e *echo.Group) {
	transRepository := repository.RepositoryTransaction(mysql.DB)
	h := handlers.HandleTransaction(transRepository)
	e.GET("/transactions", h.FindTransactions)
	e.POST("/transaction", middleware.Auth(middleware.UploadFile(h.CreateTransaction)))
	e.GET("/transaction/:id", h.GetTransaction)
	e.PATCH("/transaction-update/:id", middleware.Auth(h.UpdateTransaction))
	e.DELETE("/transaction-delete/:id", middleware.Auth(h.DeleteTransaction))
}
