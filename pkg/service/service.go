package service

import (
	to_do "github.com/tiquiet/to-do"
	"github.com/tiquiet/to-do/pkg/repository"
)

type Authorization interface {
	CreateUser(user to_do.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list to_do.CreateListItem) (int, error)
	GetAll(userId int) ([]to_do.TodoList, error)
	GetById(userId, listId int) (to_do.TodoList, error)
	DeleteById(userId, listId int) error
	Update(userId, listId int, input to_do.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item to_do.CreateTodoItem) (int, error)
	GetAll(userId, listId int) ([]to_do.TodoItem, error)
	GetItemById(userId, itemId int) (to_do.TodoItem, error)
	DeleteItem(userId, itemId int) error
	UpdateItem(userId, itemId int, input to_do.UpdateListItemInput) error
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
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
