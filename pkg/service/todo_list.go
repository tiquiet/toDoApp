package service

import (
	to_do "github.com/tiquiet/to-do"
	"github.com/tiquiet/to-do/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list to_do.CreateListItem) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]to_do.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listId int) (to_do.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *TodoListService) DeleteById(userId, listId int) error {
	return s.repo.DeleteById(userId, listId)
}

func (s *TodoListService) Update(userId, listId int, input to_do.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)
}
