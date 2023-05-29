package middleware

import (
	resultdto "dumbmerch/dto/result"
	jwtToken "dumbmerch/pkg/jwt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Auth(n echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return c.JSON(http.StatusUnauthorized, resultdto.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: "Unauthorized or invalid token",
			})
		}
		token = strings.Split(token, " ")[1]
		claims, err := jwtToken.DecodeToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, Result{
				Code:    http.StatusUnauthorized,
				Message: "Token unauthorized",
			})
		}
		c.Set("userLogin", claims)
		return n(c)
	}
}
