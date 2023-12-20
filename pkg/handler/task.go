package handler

import (
	"net/http"
	"strconv"

	"github.com/ciiska5/todo-app/entities"
	"github.com/labstack/echo"
)

// получение всех задач списка
func (h *Handler) getAllTasksOfList(c echo.Context) error {
	userId, err := getUserId(c)
	if err != nil {
		return err
	}

	strListId := c.Param("id")

	listId, err := strconv.Atoi(strListId)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid list_id param")
	}

	tasks, err := h.services.ToDoTask.GetAllTasks(userId, listId)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, tasks)
}

// получение конкретной задачи спсика
func (h *Handler) getTaskById(c echo.Context) error {
	userId, err := getUserId(c)
	if err != nil {
		return err
	}

	strTaskId := c.Param("id")

	taskId, err := strconv.Atoi(strTaskId)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid task_id param")
	}

	task, err := h.services.ToDoTask.GetTaskById(userId, taskId)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, task)
}

// добавление задачи в список
func (h *Handler) addTaskToList(c echo.Context) error {
	userId, err := getUserId(c)
	if err != nil {
		return err
	}

	strListId := c.Param("id")

	listId, err := strconv.Atoi(strListId)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid list_id param")
	}

	var input entities.TodoListTask

	if err := c.Bind(&input); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	taskId, err := h.services.ToDoTask.AddTask(userId, listId, input)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"taskId":  taskId,
		"listId":  listId,
		"message": "task added successfully",
	})
}

// обновление задачи спсика
func (h *Handler) updateTaskById(c echo.Context) error {
	userId, err := getUserId(c)
	if err != nil {
		return err
	}

	strTaskId := c.Param("id")

	taskId, err := strconv.Atoi(strTaskId)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid task_id param")
	}

	var input entities.UpdateTaskInput

	if err := c.Bind(&input); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := h.services.ToDoTask.UpdateTaskById(userId, taskId, input); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	message := "task with id = " + strTaskId + " updated successfully"
	return c.JSON(http.StatusOK, map[string]string{
		"message:": message,
	})
}

// удаление задачи спсика
func (h *Handler) deleteTaskById(c echo.Context) error {
	userId, err := getUserId(c)
	if err != nil {
		return err
	}

	strTaskId := c.Param("id")

	taskId, err := strconv.Atoi(strTaskId)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid task_id param")
	}

	listId, err := h.services.ToDoTask.DeleteTaskById(userId, taskId)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	message := "task with id = " + strTaskId + " deleted successfully from list with id = " + strconv.Itoa(listId)
	return c.JSON(http.StatusOK, map[string]string{
		"message:": message,
	})
}
