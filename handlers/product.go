package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Product struct {
	Id          string `json:"id"`
	NameProduct string `json:"name"`
	Price       int    `json:"price"`
}

var Products = []Product{
	{
		Id:          "1",
		NameProduct: "Mouse",
		Price:       1200,
	},
	{
		Id:          "1",
		NameProduct: "Keyboard",
		Price:       2000,
	},
	{
		Id:          "1",
		NameProduct: "Hetsead",
		Price:       1500,
	},
}

func FindProduct(c echo.Context) error {
	c.Response().Header().Set("Content-type", "application/json")
	c.Response().WriteHeader(http.StatusOK)

	return json.NewEncoder(c.Response()).Encode(Products)
}
