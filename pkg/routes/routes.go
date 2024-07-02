package routes

import (
	"github.com/anjush-bhargavan/todo-api/config"
	"github.com/anjush-bhargavan/todo-api/pkg/handler"
	"github.com/anjush-bhargavan/todo-api/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, todoHndlr *handler.TodoHandler,userHndlr *handler.UserHandler,cnfg *config.Config) {

	v1 := router.Group("/api/v1")
	{
		v1.POST("/signup",userHndlr.UserSignUpHandler)
		v1.POST("/login",userHndlr.UserLoginHandler)
	}

	user := v1.Group("user")
	user.Use(middleware.Authorization(cnfg.SECRETKEY))
	{
		user.POST("/todos", todoHndlr.CreateTodoHandler)
		user.GET("/todos/:id", todoHndlr.GetTodoByIDHandler)
		user.PATCH("/todos/:id", todoHndlr.ComleteTodoByIDHandler)
		user.PUT("/todos", todoHndlr.UpdateTodoHandler)
		user.DELETE("/todos/:id", todoHndlr.DeleteTodoHandler)
		user.GET("/todos", todoHndlr.ListTodosHandler)
	}
}
