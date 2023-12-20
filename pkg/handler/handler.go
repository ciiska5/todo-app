package handler

import (
	"github.com/ciiska5/todo-app/pkg/service"
	"github.com/labstack/echo"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *echo.Echo {
	router := echo.New()

	//Группа эндпоинтов для регистрации и авторизации
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	//Группы эндпоинтов для работы со списками и их задачами
	api := router.Group("/api", h.userIdentity)

	{
		lists := api.Group("/lists")
		{
			lists.GET("/", h.getAllLists)          //получение всех списков
			lists.GET("/:id", h.getListById)       //получение конкретного списка
			lists.POST("/", h.addList)             //добавление спсика
			lists.PUT("/:id", h.updateListById)    //обновление конкретного спсика
			lists.DELETE("/:id", h.deleteListById) //удаление конкретного списка

			tasks := lists.Group("/:id/tasks")
			{
				tasks.GET("/", h.getAllTasksOfList) //получение всех задач списка
				tasks.POST("/", h.addTaskToList)    //добавление задачи в список
			}
		}

		tasks := api.Group("/tasks")
		{
			tasks.GET("/:id", h.getTaskById)       //получение конкретной задачи
			tasks.PUT("/:id", h.updateTaskById)    //обновление задачи
			tasks.DELETE("/:id", h.deleteTaskById) //удаление задачи
		}
	}
	return router
}
