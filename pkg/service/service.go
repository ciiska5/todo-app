package service

import (
	"github.com/ciiska5/todo-app/entities"
	"github.com/ciiska5/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user entities.User) (int, error)
	GenerateToken(nickname, password string) (string, error)
	ParseToken(token string) (int, error)
}

type ToDoList interface {
	AddList(userId int, list entities.TodoList) (int, error)
	GetAllLists(userId int) ([]entities.TodoList, error)
	GetListById(userId, listId int) (entities.TodoList, error)
	DeleteListById(userId, listId int) error
	UpdateListById(userId, listId int, input entities.UpdateListInput) error
}

type ToDoTask interface {
	AddTask(userId, listId int, task entities.TodoListTask) (int, error)
	GetAllTasks(userId, listId int) ([]entities.TodoListTask, error)
	GetTaskById(userId, taskId int) (entities.TodoListTask, error)
	DeleteTaskById(userId, taskId int) (int, error)
	UpdateTaskById(userId, taskId int, input entities.UpdateTaskInput) error
}

type Service struct {
	Authorization
	ToDoList
	ToDoTask
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
		ToDoList:      NewTodoListService(repository.ToDoList),
		ToDoTask:      NewTodoTaskService(repository.ToDoTask, repository.ToDoList),
	}
}
