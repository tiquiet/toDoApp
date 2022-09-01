package repository

import (
	"github.com/jmoiron/sqlx"
	to_do "github.com/tiquiet/to-do"
)

type AuthorizationRep interface {
	CreateUser(user to_do.User) (int, error)
	GetUser(username, password string) (to_do.User, error)
}

type TodoListRep interface {
	Create(userId int, list to_do.CreateListItem) (int, error)
	GetAll(userId int) ([]to_do.TodoList, error)
	GetById(userId, listId int) (to_do.TodoList, error)
	DeleteById(userId, listId int) error
	Update(userId, listId int, input to_do.UpdateListInput) error
}

type TodoItemRep interface {
	Create(listId int, item to_do.CreateTodoItem) (int, error)
	GetAll(userId, listId int) ([]to_do.TodoItem, error)
	GetItemById(userId, itemId int) (to_do.TodoItem, error)
	DeleteItem(userId, itemId int) error
	UpdateItem(userId, itemId int, input to_do.UpdateListItemInput) error
}

type Repository struct {
	AuthorizationRep
	TodoListRep
	TodoItemRep
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AuthorizationRep: NewAutPostgres(db),
		TodoListRep:      NewTodoListPostgres(db),
		TodoItemRep:      NewTodoItemPostgres(db),
	}
}
