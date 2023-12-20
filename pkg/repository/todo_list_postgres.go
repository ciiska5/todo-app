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

type ToDoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *ToDoListPostgres {
	return &ToDoListPostgres{
		db: db,
	}
}

// добавляет новый список
func (r *ToDoListPostgres) AddList(userId int, list entities.TodoList) (int, error) {
	tx, err := r.db.Begin() // создание транзакции

	if err != nil {
		return 0, err
	}

	var id int
	addListQuery := fmt.Sprintf("INSERT INTO %s (title, description, created_at) VALUES ($1, $2, $3) RETURNING id", listsTable)
	row := tx.QueryRow(addListQuery, list.Title, list.Description, time.Now())
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListsQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListsQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

// получает все списки пользователя
func (r *ToDoListPostgres) GetAllLists(userId int) ([]entities.TodoList, error) {
	var lists []entities.TodoList

	getAllListsQuery := fmt.Sprintf("SELECT lt.id, lt.title, lt.description, lt.created_at "+
		"FROM %s AS lt INNER JOIN %s AS ult "+
		"ON lt.id = ult.list_id "+
		"WHERE ult.user_id = $1", listsTable, usersListsTable)

	err := r.db.Select(&lists, getAllListsQuery, userId)

	return lists, err
}

// получает список по id
func (r *ToDoListPostgres) GetListById(userId, listId int) (entities.TodoList, error) {
	var list entities.TodoList

	getListByIdQuery := fmt.Sprintf("SELECT lt.id, lt.title, lt.description, lt.created_at "+
		"FROM %s AS lt INNER JOIN %s AS ult "+
		"ON lt.id = ult.list_id "+
		"WHERE ult.user_id = $1 AND ult.list_id = $2", listsTable, usersListsTable)

	err := r.db.Get(&list, getListByIdQuery, userId, listId)

	return list, err
}

// удаляет список по id
func (r *ToDoListPostgres) DeleteListById(userId, listId int) error {
	deleteListByIdQuery := fmt.Sprintf("DELETE FROM %s AS lt USING %s AS ult "+
		"WHERE lt.id = ult.list_id AND "+
		"ult.user_id = $1 AND ult.list_id = $2", listsTable, usersListsTable)

	_, err := r.db.Exec(deleteListByIdQuery, userId, listId)

	return err
}

// обновляет список по id
func (r *ToDoListPostgres) UpdateListById(userId, listId int, input entities.UpdateListInput) error {
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

	setQuery := strings.Join(setValue, ", ") //объединение значений слайса в одну строку

	/*три варианта после SET:
	  1) SET title=$1
	  2) SET description=$1
	  3) SET title=$1, description=$2*/
	updateQuery := fmt.Sprintf("UPDATE %s AS lt SET %s FROM %s AS ult "+
		"WHERE lt.id = ult.list_id AND ult.list_id = $%d AND ult.user_id = $%d",
		listsTable, setQuery, usersListsTable, argId, argId+1)

	args = append(args, listId, userId)

	//смотрим получившийся запрос и используемые в нем аргументы
	log.Printf("updateQuery: %s", updateQuery)
	log.Printf("args: %s", args)

	_, err := r.db.Exec(updateQuery, args...)

	return err
}
