package service

import (
	todo "learn-rest-api.go"
	"learn-rest-api.go/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}
func (service *TodoItemService) Create(userId int, listId int, item todo.TodoItem) (int, error) {
	_, err := service.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}
	return service.repo.Create(userId, listId, item)
}

func (service *TodoItemService) GetAll(userId int, listId int) ([]todo.TodoItem, error) {
	return service.repo.GetAll(userId, listId)
}

func (service *TodoItemService) GetById(userId int, itemId int) (todo.TodoItem, error) {
	return service.repo.GetById(userId, itemId)
}

func (service *TodoItemService) Update(userId int, itemId int, input todo.UpdateItemInput) error {
	return service.repo.Update(userId, itemId, input)
}

func (service *TodoItemService) Delete(userId int, itemId int) error {
	return service.repo.Delete(userId, itemId)
}
