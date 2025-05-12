package handler

import (
	"auth-service/internal/dto"
	"auth-service/internal/exception"
	responseDto "auth-service/internal/handler/dto"
	"auth-service/internal/service"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (handler *UserHandler) ChangePassword(c *gin.Context) {
	controller := responseDto.Gin{C: c}

	var req dto.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(exception.NewAppError("Invalid request format", http.StatusBadRequest))
		return
	}

	var email, exist = c.Get("username")
	if exist != true {
		c.Error(exception.NewAppError("Не удалось распознать имя пользователя", http.StatusForbidden))
		return
	}

	response, err := handler.userService.ChangePassword(email.(string), &req)
	if err != nil {
		if errors.Is(err, exception.InvalidEmail) {
			c.Error(exception.NewAppError("user not found", http.StatusNotFound))
			return
		}
		if errors.Is(err, exception.InvalidPassword) {
			c.Error(exception.NewAppError("old password doesn't match", http.StatusForbidden))
			return
		}
		c.Error(exception.NewAppError(err.Error(), http.StatusInternalServerError))
		return
	}
	controller.SuccessResponse(http.StatusOK, response)
}
