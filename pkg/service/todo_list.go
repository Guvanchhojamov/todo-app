package service

import (
	"github.com/guvanchhojamov/app-todo/pkg/model"
	"github.com/guvanchhojamov/app-todo/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list model.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAllList(userId int) ([]model.TodoList, error) {
	return s.repo.GetAllList(userId)
}
func (s *TodoListService) GetListById(userId, listId int) (model.TodoList, error) {
	return s.repo.GetListById(userId, listId)
}

func (s *TodoListService) DeleteList(userId int, listId int) error {
	return s.repo.DeleteList(userId, listId)
}

func (s *TodoListService) UpdateList(input model.UpdateListInput, listId int, userId int) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateList(input, listId, userId)
}
