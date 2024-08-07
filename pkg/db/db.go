package db

import (
	"fmt"
	"github.com/anjush-bhargavan/todo-api/config"
	"github.com/gocql/gocql"
)

// ConnectScylla establishes a connection to the ScyllaDB and returns the session
func ConnectScylla(config *config.Config) (*gocql.Session, error) {
	cluster := gocql.NewCluster(config.DBHost)
	cluster.Keyspace = config.DBKeyspace
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, fmt.Errorf("error creating ScyllaDB session: %v", err)
	}

	err = createTable(session)
	if err != nil {
		session.Close()
		return nil, fmt.Errorf("error creating table: %v", err)
	}

	return session, nil
}

// createTable creates the todos table in ScyllaDB
func createTable(session *gocql.Session) error {
    todoTableQuery := `
    CREATE TABLE IF NOT EXISTS todos (
        id UUID,
        user_id UUID,
        title TEXT,
        description TEXT,
        status TEXT,
        created TIMESTAMP,
        updated TIMESTAMP,
        PRIMARY KEY (user_id, id)
    );
    `
    userTableQuery := `
    CREATE TABLE IF NOT EXISTS users (
        id UUID,
        username TEXT,
        email TEXT,
        password TEXT,
		PRIMARY KEY (email, id)
    );
    `
    
    err := session.Query(todoTableQuery).Exec()
    if err != nil {
        return fmt.Errorf("error creating todos table: %v", err)
    }

    err = session.Query(userTableQuery).Exec()
    if err != nil {
        return fmt.Errorf("error creating users table: %v", err)
    }

    return nil
}