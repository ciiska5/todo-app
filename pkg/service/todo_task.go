package service

import (
	"github.com/ciiska5/todo-app/entities"
	"github.com/ciiska5/todo-app/pkg/repository"
)

/* Сервис для работы с задачами списка */

type ToDoTaskService struct {
	repo     repository.ToDoTask
	listRepo repository.ToDoList
}

func NewTodoTaskService(repo repository.ToDoTask, listRepo repository.ToDoList) *ToDoTaskService {
	return &ToDoTaskService{
		repo:     repo,
		listRepo: listRepo,
	}
}

// добавляет новую задачу
func (s *ToDoTaskService) AddTask(userId, listId int, task entities.TodoListTask) (int, error) {
	//проверка существования списка и принадлажности списка пользователю
	_, err := s.listRepo.GetListById(userId, listId)
	if err != nil {
		return 0, err
	}

	return s.repo.AddTask(listId, task)
}

// получает все задачи списка
func (s *ToDoTaskService) GetAllTasks(userId, listId int) ([]entities.TodoListTask, error) {
	return s.repo.GetAllTasks(userId, listId)
}

// получает задачу по id
func (s *ToDoTaskService) GetTaskById(userId, taskId int) (entities.TodoListTask, error) {
	return s.repo.GetTaskById(userId, taskId)
}

// удаляет задачу по id
func (s *ToDoTaskService) DeleteTaskById(userId, taskId int) (int, error) {
	return s.repo.DeleteTaskById(userId, taskId)
}

// обновление задачи по id
func (s *ToDoTaskService) UpdateTaskById(userId, taskId int, input entities.UpdateTaskInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateTaskById(userId, taskId, input)
}
