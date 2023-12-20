package repository

/*Логика подключения БД*/

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// константы для таблиц из БД
const (
	usersListsTable = "users_lists"
	usersTable      = "users"
	listsTasksTable = "lists_tasks"
	listsTable      = "lists"
	tasksTable      = "tasks"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgesDB(config Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.DBName, config.Password, config.SSLMode))

	if err != nil {
		return nil, err
	}

	//проверка подключения
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
