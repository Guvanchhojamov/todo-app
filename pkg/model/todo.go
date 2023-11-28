package model

import "errors"

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title"  db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type TodoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil || *i.Title == "" {
		return errors.New("update fields has no values")
	}
	return nil
}

func (it UpdateItemInput) ValidateItemInput() error {
	if it.Title == nil && it.Description == nil || *it.Title == "" {
		return errors.New("update item fields has no values")
	}
	return nil
}
