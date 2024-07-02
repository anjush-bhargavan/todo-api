package service

import (
	"errors"
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
	todo.Status = "pending"
	todo.Created = time.Now()
	todo.Updated = time.Now()
	return s.Repo.CreateTodo(todo)
}

func (s *TodoService) GetTodoByIDSvc(id string, userID gocql.UUID) (*models.Todo, error) {

	todo, err := s.Repo.GetTodoByID(id)
	if err != nil {
		return nil, err
	}
	if todo.UserID != userID {
		return nil, errors.New("unauthorised user")
	}

	return todo, nil
}

func (s *TodoService) UpdateTodoSvc(todo *models.Todo) error {
	todo.Updated = time.Now()
	currentTodo, err := s.Repo.GetTodoByID(todo.ID.String())
	if err != nil {
		return err
	}
	if todo.UserID != currentTodo.UserID {
		return errors.New("unauthorised user")
	}

	return s.Repo.UpdateTodo(todo)
}

func (s *TodoService) DeleteTodoSvc(id string, userID gocql.UUID) error {

	todo, err := s.Repo.GetTodoByID(id)
	if err != nil {
		return err
	}
	if todo.UserID != userID {
		return errors.New("unauthorised user")
	}

	return s.Repo.DeleteTodo(id)
}

func (s *TodoService) ListTodosSvc(limit, offset int, userID, status string) ([]*models.Todo, error) {
	return s.Repo.ListTodos(limit, offset, userID, status)
}

func (s *TodoService) CompleteTodoSvc(id string, userID gocql.UUID) (*models.Todo, error) {
	todo, err := s.Repo.GetTodoByID(id)
	if err != nil {
		return nil, err
	}
	
	if todo.UserID != userID {
		return nil,errors.New("unauthorised user")
	}

	todo.Updated = time.Now()
	todo.Status = "completed"

	return todo,s.Repo.UpdateTodo(todo)
}