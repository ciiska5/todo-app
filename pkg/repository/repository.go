package repository

import (
	"github.com/ciiska5/todo-app/entities"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user entities.User) (int, error)
	GetUser(nickname, password string) (entities.User, error)
}

type ToDoList interface {
	AddList(userId int, list entities.TodoList) (int, error)
	GetAllLists(userId int) ([]entities.TodoList, error)
	GetListById(userId, listId int) (entities.TodoList, error)
	DeleteListById(userId, listId int) error
	UpdateListById(userId, listId int, input entities.UpdateListInput) error
}

type ToDoTask interface {
	AddTask(listId int, task entities.TodoListTask) (int, error)
	GetAllTasks(userId, listId int) ([]entities.TodoListTask, error)
	GetTaskById(userId, taskId int) (entities.TodoListTask, error)
	DeleteTaskById(userId, taskId int) (int, error)
	UpdateTaskById(userId, taskId int, input entities.UpdateTaskInput) error
}

type Repository struct {
	Authorization
	ToDoList
	ToDoTask
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPodtgres(db),
		ToDoList:      NewTodoListPostgres(db),
		ToDoTask:      NewTodoTaskPostgres(db),
	}
}
