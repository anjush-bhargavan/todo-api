package handler

import (
	"net/http"
	"strconv"

	"github.com/anjush-bhargavan/todo-api/pkg/models"
	"github.com/anjush-bhargavan/todo-api/pkg/service/interfaces"
	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	Service interfaces.TodoServiceInter
}

func NewTodoHandler(service interfaces.TodoServiceInter) *TodoHandler {
	return &TodoHandler{Service: service}
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.CreateTodoSvc(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, todo)
}

func (h *TodoHandler) GetTodoByID(c *gin.Context) {
	todoID := c.Param("id")
	todo, err := h.Service.GetTodoByIDSvc(todoID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.UpdateTodoSvc(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	todoID := c.Param("id")
	if err := h.Service.DeleteTodoSvc(todoID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}

func (h *TodoHandler) ListTodos(c *gin.Context) {
	userID := c.Query("user_id")
	status := c.Query("status")
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	todos, err := h.Service.ListTodoSvc(userID, limit, offset, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}
