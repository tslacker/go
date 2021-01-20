package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type a struct {
	X  int    `json:"x"`
	Y  string `json:"y"`
	Id int    `json:"id"`
}

type Person struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"hello": "world"})
	})

	e.POST("/aaaa", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"hello": "Go"})
	})
	e.POST("/bbbb", func(c echo.Context) error {
		var b a
		b.X = 11
		b.Y = "GOLANG"
		b.Id = 4444
		return c.JSON(http.StatusOK, b)
	})
	e.Logger.Fatal(e.Start(":1313"))
}
