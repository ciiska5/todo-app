package handler

/*Прослойка парсинга JWT и предоставления доствупа к группе эндпоинтов api */

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

const (
	authorizationHeader = echo.HeaderAuthorization
	userCtx             = "userId"
)

func (h *Handler) userIdentity(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get(authorizationHeader)
		if header == "" {
			return c.String(http.StatusUnauthorized, "empty auth header")
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			return c.String(http.StatusUnauthorized, "invalid auth header")
		}

		userId, err := h.services.Authorization.ParseToken(headerParts[1])
		if err != nil {
			return c.String(http.StatusUnauthorized, err.Error())
		}

		c.Set(userCtx, userId)
		return next(c)
	}
}

// преобразует тип данных userId из итерфейса в int
func getUserId(c echo.Context) (int, error) {
	id := c.Get(userCtx)
	if id == nil {
		return 0, c.String(http.StatusInternalServerError, "user id not found")
	}

	idInt, ok := id.(int)

	if !ok {
		return 0, c.String(http.StatusInternalServerError, "user id not found")
	}

	return idInt, nil
}
