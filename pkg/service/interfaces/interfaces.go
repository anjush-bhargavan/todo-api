package interfaces

import (
	"github.com/anjush-bhargavan/todo-api/pkg/models"
	"github.com/gocql/gocql"
)

type TodoServiceInter interface {
	CreateTodoSvc(todo *models.Todo) error
	GetTodoByIDSvc(id string, userID gocql.UUID) (*models.Todo, error)
	UpdateTodoSvc(todo *models.Todo) error
    CompleteTodoSvc(id string, userID gocql.UUID) (*models.Todo, error)
	DeleteTodoSvc(id string, userID gocql.UUID) error
	ListTodosSvc(limit, offset int, userID, status string) ([]*models.Todo, error)
   
}

type UserServiceInter interface {
	UserSignUpSvc(user *models.User) error
	UserLoginSvc(login *models.Login) (string, error)
}
