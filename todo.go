package main

import (
     "net/http"

     "github.com/labstack/echo"
     "github.com/labstack/echo/middleware"
)

func main() {

    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/", func(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
    })

    e.GET("/tasks", func(c echo.Context) error {
	return c.JSON(200, "GET Tasks")
    })

    e.PUT("/tasks", func(c echo.Context) error {
	return c.JSON(200, "PUT Tasks")
    })

    e.DELETE("/tasks/:id", func(c echo.Context) error {
	return c.JSON(200, "DELETE Task " + c.Param("id"))
    })

		// Start server
    e.Logger.Fatal(e.Start(":1323"))
}
