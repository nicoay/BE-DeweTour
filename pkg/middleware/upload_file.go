package middleware

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UploadFile(n echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("image")
		if err != nil {
			return c.JSON(http.StatusBadRequest, "failed get images")
		}
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, "failed open images")
		}
		defer src.Close()

		tempFile, err := ioutil.TempFile("uploads", "image-*.png")
		if err != nil {
			return c.JSON(http.StatusBadRequest, "failed upload file")
		}
		defer tempFile.Close()

		if _, err = io.Copy(tempFile, src); err != nil {
			return c.JSON(http.StatusBadRequest, "failed copy images")
		}

		data := tempFile.Name()
		filename := data[8:]

		c.Set("dataFile", filename)
		return n(c)
	}
}
