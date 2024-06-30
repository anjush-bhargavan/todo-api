package routes

import (
	"github.com/anjush-bhargavan/todo-api/pkg/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, hndlr *handler.TodoHandler) {

	v1 := router.Group("/api/v1")
	{
		v1.POST("/todos", hndlr.CreateTodo)
		v1.GET("/todos/:id", hndlr.GetTodoByID)
		v1.PUT("/todos/:id", hndlr.UpdateTodo)
		v1.DELETE("/todos/:id", hndlr.DeleteTodo)
		v1.GET("/todos", hndlr.ListTodos)
	}
}
