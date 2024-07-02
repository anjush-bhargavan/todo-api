package models

import (
	"time"

	"github.com/gocql/gocql"
)

// Todo struct represents the todo list data
type Todo struct {
	ID          gocql.UUID `json:"id"`
	UserID      gocql.UUID `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	Created     time.Time  `json:"created"`
	Updated     time.Time  `json:"updated"`
}

// User struct represents the user data
type User struct {
    ID       gocql.UUID `json:"id"`
    Username string     `json:"username"`
    Email    string     `json:"email"`
    Password string     `json:"password"`
}

// Login struct represents the user login data
type Login struct {
    Email    string     `json:"email"`
    Password string     `json:"password"`
}