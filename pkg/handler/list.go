package handler

import (
	"net/http"
	"strconv"

	"github.com/ciiska5/todo-app/entities"
	"github.com/labstack/echo"
)

// получение всех списков пользователя
type getAllListsResponse struct {
	Data []entities.TodoList `json:"data"`
}

func (h *Handler) getAllLists(c echo.Context) error {
	userId, err := getUserId(c)
	if err != nil {
		return err
	}

	lists, err := h.services.ToDoList.GetAllLists(userId)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

// получение конкретного списка
func (h *Handler) getListById(c echo.Context) error {
	userId, err := getUserId(c)
	if err != nil {
		return err
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id param")
	}

	list, err := h.services.ToDoList.GetListById(userId, listId)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, list)
}

// добавление спсика
func (h *Handler) addList(c echo.Context) error {
	userId, err := getUserId(c)
	if err != nil {
		return err
	}

	var input entities.TodoList

	if err := c.Bind(&input); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	id, err := h.services.ToDoList.AddList(userId, input)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// обновление конкретного спсика
func (h *Handler) updateListById(c echo.Context) error {
	userId, err := getUserId(c)
	if err != nil {
		return err
	}

	strListId := c.Param("id")

	listId, err := strconv.Atoi(strListId)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id param")
	}

	var input entities.UpdateListInput

	if err := c.Bind(&input); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := h.services.ToDoList.UpdateListById(userId, listId, input); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	message := "list with id = " + strListId + " updated successfully"
	return c.JSON(http.StatusOK, map[string]string{
		"message:": message,
	})
}

// удаление конкретного списка
func (h *Handler) deleteListById(c echo.Context) error {
	userId, err := getUserId(c)
	if err != nil {
		return err
	}

	strListId := c.Param("id")

	listId, err := strconv.Atoi(strListId)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id param")
	}

	err = h.services.ToDoList.DeleteListById(userId, listId)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	message := "list with id = " + strListId + " deleted successfully"
	return c.JSON(http.StatusOK, map[string]string{
		"message:": message,
	})
}
