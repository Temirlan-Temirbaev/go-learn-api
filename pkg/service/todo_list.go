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

func (service *TodoListService) GetAll(userId int) ([]todo.TodoList, error) {
	return service.repo.GetAll(userId)
}

func (service *TodoListService) GetById(userId int, todoId int) (todo.TodoList, error) {
	return service.repo.GetById(userId, todoId)
}
