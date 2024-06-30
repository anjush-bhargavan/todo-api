package di

import (
	"log"

	"github.com/anjush-bhargavan/todo-api/config"
	"github.com/anjush-bhargavan/todo-api/pkg/db"
	"github.com/anjush-bhargavan/todo-api/pkg/handler"
	"github.com/anjush-bhargavan/todo-api/pkg/repo"
	"github.com/anjush-bhargavan/todo-api/pkg/routes"
	"github.com/anjush-bhargavan/todo-api/pkg/server"
	"github.com/anjush-bhargavan/todo-api/pkg/service"
)

func Init() {
	cnfg := config.LoadConfig()

	session, err := db.ConnectScylla(cnfg)
	if err != nil {
		log.Fatalf("failed to connect to ScyllaDB")
	}

	todoRepo := repo.NewTodoRepository(session)
	todoSvc := service.NewTodoService(todoRepo)
	todoHandler := handler.NewTodoHandler(todoSvc)

	server := server.NewServer()

	routes.RegisterRoutes(server.R,todoHandler)

	server.StartServer(cnfg.PORT)
}
