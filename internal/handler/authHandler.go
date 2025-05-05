package handler

import (
	"auth-service/internal/exception"
	"net/http"

	"auth-service/internal/domain"
	"auth-service/internal/handler/dto"
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
	controller := dto.Gin{C: c}

	var req domain.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(exception.NewAppError("Invalid request format", http.StatusBadRequest))
		return
	}

	response, err := h.userService.Register(&req)

	if err != nil {
		c.Error(err)
		return
	}

	controller.SuccessResponse(http.StatusCreated, response)
}

func (h *AuthHandler) Login(c *gin.Context) {
	controller := dto.Gin{C: c}

	var req domain.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(exception.NewAppError("Invalid request format", http.StatusBadRequest))
		return
	}

	response, err := h.userService.Login(&req)
	if err != nil {
		c.Error(err)
		return
	}
	controller.SuccessResponse(http.StatusOK, response)
}

func (h *AuthHandler) RefreshTokens(c *gin.Context) {
	controller := dto.Gin{C: c}

	var req domain.TokenCoupleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(exception.NewAppError("Invalid request format", http.StatusBadRequest))
		return
	}

	response, err := h.userService.RefreshTokens(&req)
	if err != nil {
		c.Error(err)
		return
	}
	controller.SuccessResponse(http.StatusOK, response)
}
