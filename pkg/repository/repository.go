package repository

import (
	"github.com/guvanchhojamov/app-todo/pkg/model"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUserFromDB(username, password string) (model.User, error)
}

type TodoList interface {
	Create(userId int, list model.TodoList) (int, error)
	GetAllList(userId int) ([]model.TodoList, error)
	GetListById(userId int, listId int) (model.TodoList, error)
	DeleteList(userId int, listId int) error
	UpdateList(input model.UpdateListInput, listId int, userId int) error
}

type TodoItem interface {
	CreateItem(input model.TodoItem, listId int) (int, error)
	GetUserAllItems(userId, listId int) ([]model.TodoItem, error)
	GetItemById(userId, listId int) (model.TodoItem, error)
	DeleteItem(userId, listId int) error
	UpdateItem(input model.UpdateItemInput, userId, listId int) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
