package service

import (
	"time"

	"github.com/anjush-bhargavan/todo-api/pkg/models"
	inter "github.com/anjush-bhargavan/todo-api/pkg/repo/interfaces"
	"github.com/anjush-bhargavan/todo-api/pkg/service/interfaces"
	"github.com/gocql/gocql"
)

type TodoService struct {
	Repo inter.TodoRepositoryInter
}

func NewTodoService(repo inter.TodoRepositoryInter) interfaces.TodoServiceInter {
	return &TodoService{Repo: repo}
}

func (s *TodoService) CreateTodoSvc(todo *models.Todo) error {
	todo.ID = gocql.TimeUUID()
	todo.Created = time.Now()
	todo.Updated = time.Now()
	return s.Repo.CreateTodo(todo)
}

func (s *TodoService) GetTodoByIDSvc(id string) (*models.Todo, error) {
	return s.Repo.GetTodoByID(id)
}

func (s *TodoService) UpdateTodoSvc(todo *models.Todo) error {
	todo.Updated = time.Now()
	return s.Repo.UpdateTodo(todo)
}

func (s *TodoService) DeleteTodoSvc(id string) error {
	return s.Repo.DeleteTodo(id)
}

func (s *TodoService) ListTodoSvc(userID string, limit, offset int, status string) ([]*models.Todo, error) {
	return s.Repo.ListTodo(userID, limit, offset, status)
}
