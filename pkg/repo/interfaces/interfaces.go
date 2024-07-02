package interfaces

import (
	"github.com/anjush-bhargavan/todo-api/pkg/models"
)

type TodoRepositoryInter interface {
	CreateTodo(todo *models.Todo) error
	GetTodoByID(id string) (*models.Todo, error)
	UpdateTodo(todo *models.Todo) error
	DeleteTodo(id string) error
	ListTodos(limit, offset int, userID, status string) ([]*models.Todo, error)
}

type UserRepositoryInter interface {
	CreateUser(user *models.User) error
	GetUserByID(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}
