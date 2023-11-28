package repository

import (
	"fmt"
	"github.com/guvanchhojamov/app-todo/pkg/model"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (authPostgres *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name,username,password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	row := authPostgres.db.QueryRow(query, user.Name, user.Username, user.Password)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (authPostgres *AuthPostgres) GetUserFromDB(username, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := authPostgres.db.Get(&user, query, username, password)
	return user, err
}
