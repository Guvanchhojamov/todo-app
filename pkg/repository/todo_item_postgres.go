package repository

import (
	"errors"
	"fmt"
	"github.com/guvanchhojamov/app-todo/pkg/model"
	"github.com/jmoiron/sqlx"
	"strings"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) CreateItem(input model.TodoItem, listId int) (int, error) {
	trx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title,description,done) VALUES ($1,$2,$3) RETURNING id", itemsTable)
	row := r.db.QueryRow(createItemQuery, input.Title, input.Description, input.Done)

	if err = row.Scan(&itemId); err != nil {
		err := trx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}
	createListsItemsQuery := fmt.Sprintf(`INSERT INTO %s  (list_id, item_id) VALUES($1, $2) RETURNING id`, listsItemsTable)
	_, err = r.db.Exec(createListsItemsQuery, listId, itemId)
	if err != nil {
		return 0, trx.Rollback()
	}
	return itemId, trx.Commit()
}

func (r *TodoItemPostgres) GetUserAllItems(userId, listId int) ([]model.TodoItem, error) {

	var items []model.TodoItem
	getAllitemsQuery := fmt.Sprintf(`SELECT ti.id, ti.title,ti.description, ti.done FROM %s ti INNER JOIN %s li ON ti.id = li.item_id 
    										INNER JOIN %s ul ON ul.list_id = li.list_id WHERE li.list_id=$1 AND ul.user_id=$2`,
		itemsTable, listsItemsTable, usersListsTable)
	if err := r.db.Select(&items, getAllitemsQuery, listId, userId); err != nil {
		return nil, err
	}
	fmt.Println(getAllitemsQuery)
	return items, nil
}

func (r *TodoItemPostgres) GetItemById(userId, itemId int) (model.TodoItem, error) {
	var item model.TodoItem

	getItemsByIdQuery := fmt.Sprintf(`SELECT ti.id, ti.title,ti.description, ti.done FROM %s ti INNER JOIN %s li ON ti.id = li.item_id 
    										INNER JOIN %s ul ON ul.list_id = li.list_id WHERE ti.id=$1 AND ul.user_id=$2`,
		itemsTable, listsItemsTable, usersListsTable)

	if err := r.db.Get(&item, getItemsByIdQuery, itemId, userId); err != nil {
		errors.New("get items By id db get Error")
		return item, err
	}
	return item, nil
}

func (r *TodoItemPostgres) DeleteItem(userId, itemId int) error {
	delItemQuery := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul
		 								WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id =$1 AND ti.id=$2`,
		itemsTable, listsItemsTable, usersListsTable)
	_, err := r.db.Exec(delItemQuery, userId, itemId)
	return err
}

func (r *TodoItemPostgres) UpdateItem(input model.UpdateItemInput, userId, itemId int) error {
	var setValues = make([]string, 0)
	var args = make([]interface{}, 0)
	var argId = 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setString := strings.Join(setValues, ", ")
	args = append(args, userId, itemId)
	updateItemString := fmt.Sprintf(`UPDATE %s ti SET %s FROM  %s li, %s ul 
											WHERE  ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id=$%d AND ti.id=$%d`,
		itemsTable, setString, listsItemsTable, usersListsTable, argId, argId+1)

	fmt.Printf("update item query: \n", updateItemString)
	// fmt.Println(args...)
	_, err := r.db.Exec(updateItemString, args...)
	return err
}
