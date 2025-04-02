package dto

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

func (g *Gin) SuccessResponse(httpCode int, data interface{}) {
	g.C.JSON(httpCode, SuccessResponse{
		Success: true,
		Data:    data,
	})
	return
}

func (g *Gin) ErrorResponse(httpCode int, reason string, errorMessage string) {
	g.C.JSON(httpCode, ErrorResponse{
		Success: false,
		Reason:  reason,
		Message: errorMessage,
	})
	return
}
