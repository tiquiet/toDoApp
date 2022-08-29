package service

import (
	to_do "github.com/tiquiet/to-do"
	"github.com/tiquiet/to-do/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int, list to_do.CreateTodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(listId, list)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]to_do.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *TodoItemService) GetItemById(userId, itemId int) (to_do.TodoItem, error) {
	return s.repo.GetItemById(userId, itemId)
}

func (s *TodoItemService) DeleteItem(userId, itemId int) error {
	return s.repo.DeleteItem(userId, itemId)
}

func (s *TodoItemService) UpdateItem(userId, itemId int, input to_do.UpdateListItemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateItem(userId, itemId, input)
}
