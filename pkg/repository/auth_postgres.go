package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	todo "learn-rest-api.go"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (repository *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password) values ($1, $2, $3) RETURNING id", usersTable)
	row := repository.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (repository *AuthPostgres) GetUser(username string, password string) (todo.User, error) {
	var user todo.User
	query := fmt.Sprintf("SELECT id from %s WHERE username=$1 AND password=$2", usersTable)
	err := repository.db.Get(&user, query, username, password)
	return user, err
}
