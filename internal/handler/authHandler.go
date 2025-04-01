package handler

import (
	"net/http"
	"strings"

	"auth-service/internal/domain"

	"auth-service/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService *service.UserService
}

func NewAuthHandler(userService *service.UserService) *AuthHandler {
	return &AuthHandler{userService: userService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req domain.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": "false",
			"error":   "Invalid request format",
			"details": gin.H{
				"validation_errors": err.Error(),
			},
		})
		return
	}

	response, err := h.userService.Register(&req)

	if err != nil {
		status := http.StatusInternalServerError
		message := err.Error()

		c.JSON(status, gin.H{
			"success": "false",
			"error":   message,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": "true",
		"data":    response,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req domain.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": "false",
			"error":   "Invalid request format",
			"details": gin.H{
				"validation_errors": err.Error(),
			},
		})
		return
	}

	response, err := h.userService.Login(&req)
	if err != nil {
		status := http.StatusUnauthorized
		message := err.Error()

		if strings.Contains(err.Error(), "not found") {
			message = "Invalid email or password"
		}

		c.JSON(status, gin.H{
			"success": "false",
			"error":   message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "true",
		"data":    response,
	})
}
