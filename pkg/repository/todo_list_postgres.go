package repository

import (
	"fmt"
	"github.com/guvanchhojamov/app-todo/pkg/model"
	"github.com/jmoiron/sqlx"
	"strings"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userId int, list model.TodoList) (int, error) {
	trx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", listsTable)
	row := trx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		trx.Rollback()
		return 0, err
	}
	createUsersListsQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2) RETURNING id", usersListsTable)
	_, err = trx.Exec(createUsersListsQuery, userId, id)
	if err != nil {
		trx.Rollback()
		return 0, err
	}
	return id, trx.Commit()
}

func (r *TodoListPostgres) GetAllList(userId int) ([]model.TodoList, error) {
	var lists []model.TodoList

	queryString := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id=ul.list_id WHERE ul.user_id = $1", listsTable, usersListsTable)
	fmt.Println(queryString)
	err := r.db.Select(&lists, queryString, userId)
	return lists, err
}
func (r *TodoListPostgres) GetListById(userId, listId int) (model.TodoList, error) {
	var list model.TodoList

	queryString := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1 AND tl.id = $2`, listsTable, usersListsTable)
	fmt.Println(queryString)
	err := r.db.Get(&list, queryString, userId, listId)

	return list, err
}

func (r *TodoListPostgres) DeleteList(userId int, listId int) error {
	queryString := fmt.Sprintf(`DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id =$1 AND ul.list_id = $2`, listsTable, usersListsTable)
	fmt.Println(queryString)
	_, err := r.db.Exec(queryString, userId, listId)

	return err
}

func (r *TodoListPostgres) UpdateList(input model.UpdateListInput, listId int, userId int) error {
	var setValues = make([]string, 0)
	var args = make([]interface{}, 0)
	var argId = 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title) // miss pointer for check
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description) //miss pointer for check
		argId++
	}

	setString := strings.Join(setValues, ", ")
	updateQuery := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id=ul.list_id AND ul.user_id =$%d AND ul.list_id=$%d",
		listsTable, setString, usersListsTable, argId, argId+1)

	args = append(args, userId, listId)

	fmt.Printf("updated Query:\n %s \n", updateQuery)
	fmt.Printf("args:\n %s", args)
	_, err := r.db.Exec(updateQuery, args...)
	return err
}
