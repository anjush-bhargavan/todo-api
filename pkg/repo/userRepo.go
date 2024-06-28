package repo

import (
	"github.com/anjush-bhargavan/todo-api/pkg/models"
	"github.com/anjush-bhargavan/todo-api/pkg/repo/interfaces"
	"github.com/gocql/gocql"
)

type UserRepository struct {
	Session *gocql.Session
}

func NewUserRepository(session *gocql.Session)  interfaces.UserRepositoryInter {
	return &UserRepository{Session: session}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (id, username, email, password) VALUES (?, ?, ?, ?)`
	if err := r.Session.Query(query, user.ID, user.Username, user.Email, user.Password).Exec(); err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByID(id string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, email, password FROM users WHERE id = ?`
	if err := r.Session.Query(query, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
		return nil, err
	}
	return &user, nil
}
