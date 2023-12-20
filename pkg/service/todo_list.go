package service

import (
	"github.com/ciiska5/todo-app/entities"
	"github.com/ciiska5/todo-app/pkg/repository"
)

/* Сервис для работы со списками */

type ToDoListService struct {
	repo repository.ToDoList
}

func NewTodoListService(repo repository.ToDoList) *ToDoListService {
	return &ToDoListService{
		repo: repo,
	}
}

// добавляет новый список
func (s *ToDoListService) AddList(userId int, list entities.TodoList) (int, error) {
	return s.repo.AddList(userId, list)
}

// получает все списки пользователя
func (s *ToDoListService) GetAllLists(userId int) ([]entities.TodoList, error) {
	return s.repo.GetAllLists(userId)
}

// получает список по id
func (s *ToDoListService) GetListById(userId, listId int) (entities.TodoList, error) {
	return s.repo.GetListById(userId, listId)
}

// удаляет списка по id
func (s *ToDoListService) DeleteListById(userId, listId int) error {
	return s.repo.DeleteListById(userId, listId)
}

// обновляет список по id
func (s *ToDoListService) UpdateListById(userId, listId int, input entities.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateListById(userId, listId, input)
}
