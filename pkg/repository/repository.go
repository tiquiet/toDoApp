package repository

import (
	"github.com/jmoiron/sqlx"
	to_do "github.com/tiquiet/to-do"
)

type Authorization interface {
	CreateUser(user to_do.User) (int, error)
	GetUser(username, password string) (to_do.User, error)
}

type TodoList interface {
	Create(userId int, list to_do.TodoList) (int, error)
	GetAll(userId int) ([]to_do.TodoList, error)
	GetById(userId, listId int) (to_do.TodoList, error)
	DeleteById(userId, listId int) error
	Update(userId, listId int, input to_do.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item to_do.TodoItem) (int, error)
	GetAll(userId, listId int) ([]to_do.TodoItem, error)
	GetItemById(userId, itemId int) (to_do.TodoItem, error)
	DeleteItem(userId, itemId int) error
	UpdateItem(userId, itemId int, input to_do.UpdateListItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAutPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
