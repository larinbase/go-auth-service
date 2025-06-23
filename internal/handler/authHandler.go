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
	userService     *service.UserService
	keycloakService *service.KeycloakService
}

func NewAuthHandler(userService *service.UserService, keycloakService *service.KeycloakService) *AuthHandler {
	return &AuthHandler{userService: userService, keycloakService: keycloakService}
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

func (h *AuthHandler) LoginV2(c *gin.Context) {
	controller := dto.Gin{C: c}

	var req domain.LoginV2Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(exception.NewAppError("Invalid request format", http.StatusBadRequest))
		return
	}

	response, err := h.userService.LoginV2(&req)
	if err != nil {
		c.Error(err)
		return
	}
	controller.SuccessResponse(http.StatusOK, response)
}

func (h *AuthHandler) SendCode(c *gin.Context) {
	controller := dto.Gin{C: c}

	var req domain.SendCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(exception.NewAppError("Invalid request format", http.StatusBadRequest))
		return
	}

	err := h.userService.SendEmailCode(&req)
	if err != nil {
		c.Error(err)
		return
	}
	controller.SuccessResponse(http.StatusOK, "successfully sent")
}

func (h *AuthHandler) RegisterKeyCloak(c *gin.Context) {
	controller := dto.Gin{C: c}

	var req domain.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(exception.NewAppError("Invalid request format", http.StatusBadRequest))
		return
	}

	ctx := c.Request.Context()
	err := h.keycloakService.RegisterUser(ctx, &req)

	if err != nil {
		c.Error(err)
		return
	}

	controller.SuccessResponse(http.StatusCreated, "successfully registered")
}
