package service

import (
	"github.com/guvanchhojamov/app-todo/pkg/model"
	"github.com/guvanchhojamov/app-todo/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
type TodoList interface {
	Create(userId int, list model.TodoList) (int, error)
	GetAllList(userId int) ([]model.TodoList, error)
	GetListById(userId, listId int) (model.TodoList, error)
	DeleteList(userId int, listId int) error
	UpdateList(input model.UpdateListInput, ListId int, userId int) error
}
type TodoItem interface {
	CreateItem(input model.TodoItem, listId int) (int, error)
	GetUserAllItems(userId, listId int) ([]model.TodoItem, error)
	GetItemById(userId, itemId int) (model.TodoItem, error)
	DeleteItem(userId, itemId int) error
	UpdateItem(input model.UpdateItemInput, userId, itemId int) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem),
	}
}
