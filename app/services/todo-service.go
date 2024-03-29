package services

import (
	"fmt"
	"sample-htmx-chi/app/configs"
	"sample-htmx-chi/app/models"
	"strconv"
)

type TodoService struct {
	config *configs.Config
}

func NewTodoService(config *configs.Config) *TodoService {
	var service TodoService
	service.config = config
	return &service
}

func (service *TodoService) List(q string) (*[]models.Todo, error) {
	var result []models.Todo
	service.config.Db.
		Where("lower(description) like lower(?)", fmt.Sprintf("%%%s%%", q)).
		Find(&result)
	return &result, nil
}

func (service *TodoService) Find(id string) (*models.Todo, error) {
	var result models.Todo
	k, err := strconv.Atoi(id)
	service.config.Db.Where("id = ?", k).First(&result)
	return &result, err
}

func (service *TodoService) Insert(todo models.Todo) error {
	service.config.Db.Create(&todo)
	return nil
}

func (service *TodoService) Update(todo models.Todo) error {
	service.config.Db.Where("id = ?", todo.Id).Updates(&todo)
	return nil
}

func (service *TodoService) Delete(id string) error {
	service.config.Db.Delete(&models.Todo{}, id)
	return nil
}
