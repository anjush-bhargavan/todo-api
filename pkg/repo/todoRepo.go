package repo

import (
	"github.com/anjush-bhargavan/todo-api/pkg/models"
	"github.com/anjush-bhargavan/todo-api/pkg/repo/interfaces"
	"github.com/gocql/gocql"
)

type TodoRepository struct {
	Session *gocql.Session
}

func NewTodoRepository(session *gocql.Session) interfaces.TodoRepositoryInter {
	return &TodoRepository{Session: session}
}

func (r *TodoRepository) CreateTodo(todo *models.Todo) error {
	query := `INSERT INTO todos (id, user_id, title, description, status, created, updated) VALUES (?, ?, ?, ?, ?, ?, ?)`
	if err := r.Session.Query(query, todo.ID, todo.UserID, todo.Title, todo.Description, todo.Status, todo.Created, todo.Updated).Exec(); err != nil {
		return err
	}
	return nil
}

func (r *TodoRepository) GetTodoByID(id string) (*models.Todo, error) {
	var todo models.Todo
	query := `SELECT id, user_id, title, description, status, created, updated FROM todos WHERE id = ?`
	if err := r.Session.Query(query, id).Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Description, &todo.Status, &todo.Created, &todo.Updated); err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepository) UpdateTodo(todo *models.Todo) error {
	query := `UPDATE todos SET title = ?, description = ?, status = ?, updated = ? WHERE id = ?`
	if err := r.Session.Query(query, todo.Title, todo.Description, todo.Status, todo.Updated, todo.ID).Exec(); err != nil {
		return err
	}
	return nil
}

func (r *TodoRepository) DeleteTodo(id string) error {
	query := `DELETE FROM todos WHERE id = ?`
	if err := r.Session.Query(query, id).Exec(); err != nil {
		return err
	}
	return nil
}

func (r *TodoRepository) ListTodo(userID string, limit, offset int, status string) ([]*models.Todo, error) {
	var todos []*models.Todo
	query := `SELECT id, user_id, title, description, status, created, updated FROM todos WHERE user_id = ?`
	if status != "" {
		query += ` AND status = ?`
	}
	query += ` LIMIT ? OFFSET ?`

	iter := r.Session.Query(query, userID, status, limit, offset).Iter()
	var todo models.Todo
	for iter.Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Description, &todo.Status, &todo.Created, &todo.Updated) {
		todos = append(todos, &todo)
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}
	return todos, nil
}
