package middleware

import (
	"auth-service/internal/exception"
	"auth-service/internal/handler/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		controller := dto.Gin{C: c}
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			if appErr, ok := err.(*exception.AppError); ok {
				controller.ErrorResponse(
					appErr.StatusCode,
					appErr.Message,
					err.Error())
				c.Abort()
				return
			}

			statusCode := http.StatusInternalServerError
			controller.ErrorResponse(
				statusCode,
				"internal server error",
				err.Error())

			c.Abort()
		}
	}
}
