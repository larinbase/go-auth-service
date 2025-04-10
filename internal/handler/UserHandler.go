package handler

import (
	"auth-service/internal/dto"
	"auth-service/internal/exception"
	responseDto "auth-service/internal/handler/dto"
	"auth-service/internal/service"
	"errors"
	"fmt"
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
		controller.ErrorResponse(http.StatusBadRequest, "Invalid request format", err.Error())
		return
	}

	var email, exist = c.Get("username")
	if exist != true {
		controller.ErrorResponse(http.StatusForbidden, "FORBIDDEN", fmt.Sprintf("Не удалось распознать имя пользователя"))
		return
	}

	response, err := handler.userService.ChangePassword(email.(string), &req)
	if err != nil {
		if errors.Is(err, exception.InvalidEmail) {
			controller.ErrorResponse(http.StatusNotFound, "user not found", err.Error())
			return
		}
		if errors.Is(err, exception.InvalidPassword) {
			controller.ErrorResponse(http.StatusForbidden, "old password doesn't match", err.Error())
			return
		}
		controller.ErrorResponse(http.StatusUnauthorized, err.Error(), err.Error())
		return
	}
	controller.SuccessResponse(http.StatusOK, response)
}
