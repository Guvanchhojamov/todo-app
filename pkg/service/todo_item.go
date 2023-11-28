package service

import (
	"github.com/guvanchhojamov/app-todo/pkg/model"
	"github.com/guvanchhojamov/app-todo/pkg/repository"
)

type TodoItemService struct {
	repo repository.TodoItem
}

func NewTodoItemService(repo repository.TodoItem) *TodoItemService {
	return &TodoItemService{repo: repo}
}
func (s *TodoItemService) CreateItem(input model.TodoItem, listId int) (int, error) {
	return s.repo.CreateItem(input, listId)
}

func (s *TodoItemService) GetUserAllItems(userId, listId int) ([]model.TodoItem, error) {
	return s.repo.GetUserAllItems(userId, listId)
}

func (s *TodoItemService) GetItemById(userId, itemId int) (model.TodoItem, error) {
	return s.repo.GetItemById(userId, itemId)
}

func (s *TodoItemService) DeleteItem(userId, itemId int) error {
	return s.repo.DeleteItem(userId, itemId)
}
func (s *TodoItemService) UpdateItem(input model.UpdateItemInput, userId, itemId int) error {
	if err := input.ValidateItemInput(); err != nil {
		return err
	}
	return s.repo.UpdateItem(input, userId, itemId)
}
