package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

var name = "Hello Felipe"

func main() {
	println(name)
	fmt.Print(name)

	e := echo.New()
	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(200, name)
	})
	e.GET("/v2/hello", getHello)
	e.POST("/name", createCar)
	e.Logger.Fatal(e.Start(":8080"))

}

func getHello(c echo.Context) error {
	return c.JSON(200, name)
}

func createCar(c echo.Context) error {
	return c.JSON(200, name)
}
