package repository

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ciiska5/todo-app/entities"
	"github.com/jmoiron/sqlx"
)

/* Репозиторий для работы со списками */

type ToDoTaskPostgres struct {
	db *sqlx.DB
}

func NewTodoTaskPostgres(db *sqlx.DB) *ToDoTaskPostgres {
	return &ToDoTaskPostgres{
		db: db,
	}
}

// добавляет новую задачу
func (r *ToDoTaskPostgres) AddTask(listId int, task entities.TodoListTask) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int

	addTaskQuery := fmt.Sprintf("INSERT INTO %s (title, description, created_at) VALUES ($1, $2, $3) RETURNING id", tasksTable)
	row := tx.QueryRow(addTaskQuery, task.Title, task.Description, time.Now())
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createListsTasksQuery := fmt.Sprintf("INSERT INTO %s (list_id, task_id) VALUES ($1, $2)", listsTasksTable)
	_, err = tx.Exec(createListsTasksQuery, listId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

// получает все задачи списка
func (r *ToDoTaskPostgres) GetAllTasks(userId, listId int) ([]entities.TodoListTask, error) {
	var tasks []entities.TodoListTask

	getTasksQuery := fmt.Sprintf("SELECT tt.id, tt.title, tt.description, tt.is_done, tt.created_at FROM %s AS tt "+
		"INNER JOIN %s AS ltt ON tt.id = ltt.task_id "+
		"INNER JOIN %s AS ult ON ltt.list_id = ult.list_id "+
		"WHERE ult.user_id = $1 AND ltt.list_id = $2",
		tasksTable, listsTasksTable, usersListsTable)

	if err := r.db.Select(&tasks, getTasksQuery, userId, listId); err != nil {
		return nil, err
	}

	return tasks, nil
}

// получает задачу по id
func (r *ToDoTaskPostgres) GetTaskById(userId, taskId int) (entities.TodoListTask, error) {
	var task entities.TodoListTask

	getTaskByIdQuery := fmt.Sprintf("SELECT tt.id, tt.title, tt.description, tt.is_done, tt.created_at FROM %s AS tt "+
		"INNER JOIN %s AS ltt ON tt.id = ltt.task_id "+
		"INNER JOIN %s AS ult ON ltt.list_id = ult.list_id "+
		"WHERE ult.user_id = $1 AND ltt.task_id = $2",
		tasksTable, listsTasksTable, usersListsTable)

	if err := r.db.Get(&task, getTaskByIdQuery, userId, taskId); err != nil {
		return task, err
	}

	return task, nil
}

// удаляет задачу по id
func (r *ToDoTaskPostgres) DeleteTaskById(userId, taskId int) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	deleteTaskByIdQuery := fmt.Sprintf("DELETE FROM %s AS tt USING %s AS ltt, %s AS ult "+
		"WHERE tt.id = ltt.task_id AND ltt.list_id = ult.list_id "+
		"AND ult.user_id = $1 AND tt.id = $2 RETURNING ult.list_id",
		tasksTable, listsTasksTable, usersListsTable)

	var list_id int

	row := tx.QueryRow(deleteTaskByIdQuery, userId, taskId)
	if err := row.Scan(&list_id); err != nil {
		return 0, err
	}

	return list_id, tx.Commit()
}

// обновляет задачу по id
func (r *ToDoTaskPostgres) UpdateTaskById(userId, taskId int, input entities.UpdateTaskInput) error {
	//слайсы для формирования запроса обновления в БД
	setValue := make([]string, 0)  //слайс значений
	args := make([]interface{}, 0) //слайс аргументов значений
	argId := 1

	if input.Title != nil {
		setValue = append(setValue, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValue = append(setValue, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.IsDone != nil {
		setValue = append(setValue, fmt.Sprintf("is_done=$%d", argId))
		args = append(args, *input.IsDone)
		argId++
	}

	setQuery := strings.Join(setValue, ", ") //объединение значений слайса в одну строку

	/*
		7 вариантов SET:
		  1) SET title=$1
		  2) SET description=$1
		  3) SET is_done=$1
		  4) SET title=$1, description=$2
		  5) SET title=$1, is_done=$2
		  6) SET description=$1, is_done=$2
		  7) SET description=$1, title=$2, is_done=$3
	*/
	updateQuery := fmt.Sprintf("UPDATE %s AS tt SET %s FROM %s AS ltt, %s AS ult "+
		"WHERE tt.id = ltt.task_id AND ult.list_id = ltt.list_id "+
		"AND ult.user_id = $%d AND tt.id = $%d",
		tasksTable, setQuery, listsTasksTable, usersListsTable, argId, argId+1)

	args = append(args, userId, taskId)

	//смотрим получившийся запрос и используемые в нем аргументы
	log.Printf("updateQuery: %s", updateQuery)
	log.Printf("args: %s", args)

	_, err := r.db.Exec(updateQuery, args...)

	return err
}
