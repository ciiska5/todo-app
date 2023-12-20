package repository

import (
	"fmt"

	"github.com/ciiska5/todo-app/entities"
	"github.com/jmoiron/sqlx"
)

/*Имплементация интерфейсов и работа с базой postgres*/

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPodtgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

// создание нового пользователя, возращает (id пользователя, ошибка)
func (a *AuthPostgres) CreateUser(user entities.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, nickname, password_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := a.db.QueryRow(query, user.Name, user.Nickname, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

// получение пользователя дла генерации JWT-токена
func (a *AuthPostgres) GetUser(nickname, password string) (entities.User, error) {
	var user entities.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE nickname=$1 AND password_hash=$2", usersTable)
	err := a.db.Get(&user, query, nickname, password)

	return user, err
}
