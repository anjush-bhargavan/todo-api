package handler

import (
	"net/http"

	"github.com/anjush-bhargavan/todo-api/pkg/models"
	"github.com/anjush-bhargavan/todo-api/pkg/service/interfaces"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service interfaces.UserServiceInter
}

func NewUserHandler(service interfaces.UserServiceInter) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) UserSignUpHandler(c *gin.Context) {
	var User models.User
	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in binding data",
			"Error":   err.Error()})
		return
	}
	if err := h.Service.UserSignUpSvc(&User); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": http.StatusInternalServerError,
			"Message": "error in signup service",
			"Error":   err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Status":  http.StatusCreated,
		"Message": "user signed up successfully",
		"Data":    "",
	})
}

func (h *UserHandler) UserLoginHandler(c *gin.Context) {
	var user models.Login
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in binding data",
			"Error":   err.Error()})
		return
	}

	token, err := h.Service.UserLoginSvc(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": http.StatusInternalServerError,
			"Message": "error in login service",
			"Error":   err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "user logged in successfully",
		"Data":    token,
	})
}

// c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
// 			"Message": "error while user id from context",
// 			"Error":   ""})
