package entities

import (
	"errors"
	"time"
)

// структура TodoList
type TodoList struct {
	Id          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at" binding:"required"`
}

type UsersTodoLists struct {
	UserId int
	ListId int
}

// струтура конкреткной таски TodoList
type TodoListTask struct {
	Id          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Description string    `json:"description" db:"description"`
	IsDone      bool      `json:"is_done" db:"is_done"`
	CreatedAt   time.Time `json:"created_at" db:"created_at" binding:"required"`
}

type ListsTodoTasks struct {
	ListId int
	TaskId int
}

// структура для обновления TodoList
type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

// проверяет, не равны ли nil обновляемые поля
func (uli UpdateListInput) Validate() error {
	if uli.Title == nil && uli.Description == nil {
		return errors.New("update fields do not have any values")
	}

	return nil
}

// структура для обновления TodoList
type UpdateTaskInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	IsDone      *bool   `json:"is_done"`
}

// проверяет, не равны ли nil обновляемые поля
func (uti UpdateTaskInput) Validate() error {
	if uti.Title == nil && uti.Description == nil && uti.IsDone == nil {
		return errors.New("update fields do not have any values")
	}

	return nil
}
