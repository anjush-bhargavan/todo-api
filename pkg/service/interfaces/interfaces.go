package interfaces

import "github.com/anjush-bhargavan/todo-api/pkg/models"

type TodoServiceInter interface {
    CreateTodoSvc(todo *models.Todo) error
    GetTodoByIDSvc(id string) (*models.Todo, error)
    UpdateTodoSvc(todo *models.Todo) error
    DeleteTodoSvc(id string) error
    ListTodoSvc(userID string, limit, offset int, status string) ([]*models.Todo, error)
}