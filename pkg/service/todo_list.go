package service

import (
	todo "learn-rest-api.go"
	"learn-rest-api.go/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}
func (service *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return service.repo.Create(userId, list)
}
