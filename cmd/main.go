package main

import (
	"github.com/Mehmet-Mg/go_echo_template/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	userHandler := &handler.UserHandler{}
	e.GET("/user", userHandler.HandleGetUser)

	e.Start(":8080")
}
