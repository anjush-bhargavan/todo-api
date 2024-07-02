package handler

import (
	"net/http"
	"strconv"

	"github.com/anjush-bhargavan/todo-api/pkg/models"
	"github.com/anjush-bhargavan/todo-api/pkg/service/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

type TodoHandler struct {
	Service interfaces.TodoServiceInter
}

func NewTodoHandler(service interfaces.TodoServiceInter) *TodoHandler {
	return &TodoHandler{Service: service}
}

func (h *TodoHandler) CreateTodoHandler(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in binding data",
			"Error":   err.Error()})
		return
	}

	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while user id from context",
			"Error":   ""})
		return
	}
	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while converting user id to string",
			"Error":   ""})
		return
	}

	userIDUuid, err := gocql.ParseUUID(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while converting user id to uuid",
			"Error":   err.Error()})
		return
	}

	todo.UserID = userIDUuid

	if err := h.Service.CreateTodoSvc(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in creating todo service",
			"Error":   err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Status":  http.StatusCreated,
		"Message": "todo created succefully",
		"Data":    todo,
	})
}

// GetTodoByIDHandler
func (h *TodoHandler) GetTodoByIDHandler(c *gin.Context) {
	todoID := c.Param("id")

	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while user id from context",
			"Error":   ""})
		return
	}
	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while converting user id to string",
			"Error":   ""})
		return
	}
	userIDUuid, err := gocql.ParseUUID(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while converting user id to uuid",
			"Error":   err.Error()})
		return
	}

	todo, err := h.Service.GetTodoByIDSvc(todoID, userIDUuid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "todo fetched",
		"Data":    todo,
	})
}

// UpdateTodoHandler
func (h *TodoHandler) UpdateTodoHandler(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in binding data",
			"Error":   err.Error()})
		return
	}
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while user id from context",
			"Error":   ""})
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while converting user id to string",
			"Error":   ""})
		return
	}
	userIDUuid, err := gocql.ParseUUID(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while converting user id to uuid",
			"Error":   err.Error()})
		return
	}
	todo.UserID = userIDUuid
	if err := h.Service.UpdateTodoSvc(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "todo updated succefully",
		"Data":    todo,
	})
}

// DeleteTodoHandler deletes the todo by id
func (h *TodoHandler) DeleteTodoHandler(c *gin.Context) {
	todoID := c.Param("id")

	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while user id from context",
			"Error":   ""})
		return
	}
	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while converting user id to string",
			"Error":   ""})
		return
	}
	userIDUuid, err := gocql.ParseUUID(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while converting user id to uuid",
			"Error":   err.Error()})
		return
	}

	if err := h.Service.DeleteTodoSvc(todoID, userIDUuid); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "Todo deleted successfully",
		"Data":    "",
	})
}

// ListTodosHandler lists the todos created by user
func (h *TodoHandler) ListTodosHandler(c *gin.Context) {

	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while user id from context",
			"Error":   ""})
		return
	}
	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while converting user id to string",
			"Error":   ""})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "Invalid limit value",
			"Error":   err.Error()})
		return
	}
	status := c.DefaultQuery("status", "")

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "Invalid offset value",
			"Error":   err.Error()})
		return
	}

	todos, err := h.Service.ListTodosSvc(limit, offset, userIDStr, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": http.StatusBadRequest,
			"Message": "Failed to fetch todos from service",
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "Todos fetched successfully",
		"Data":    todos,
	})
}

func (h *TodoHandler) ComleteTodoByIDHandler(c *gin.Context) {
	todoID := c.Param("id")

	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while user id from context",
			"Error":   ""})
		return
	}
	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while converting user id to string",
			"Error":   ""})
		return
	}
	userIDUuid, err := gocql.ParseUUID(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while converting user id to uuid",
			"Error":   err.Error()})
		return
	}

	todo, err := h.Service.CompleteTodoSvc(todoID, userIDUuid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not marked complete"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Message": "todo marked as completed",
		"Data":    todo,
	})
}
