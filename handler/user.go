package handler

import (
	"net/http"

	"github.com/Mehmet-Mg/go_echo_template/model"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (u UserHandler) HandleGetUser(c echo.Context) error {
	user := &model.User{
		Id:    1,
		Name:  "mgenc",
		Email: "mgenc@mail.com",
	}

	return c.JSON(http.StatusOK, user)
}
