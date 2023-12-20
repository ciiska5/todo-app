package handler

import (
	"fmt"
	"net/http"

	"github.com/ciiska5/todo-app/entities"
	"github.com/labstack/echo"
)

// Хендлер регистрации
func (h *Handler) signUp(c echo.Context) error {

	var input entities.User

	if err := c.Bind(&input); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, fmt.Sprint("User created with id = ", id))
}

// Хендлер авторизации

// структура, необходимая только при аутентификации пользователя
type signInInput struct {
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c echo.Context) error {

	var input signInInput

	if err := c.Bind(&input); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	token, err := h.services.Authorization.GenerateToken(input.Nickname, input.Password)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message:": "signed in succesfully",
		"token":    token,
	})
}
