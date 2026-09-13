package structure

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type ErrorInfo struct {
	Message string `json:"message"`
}

type Meta struct {
	Page       int `json:"page,omitempty"`
	PerPage    int `json:"per_page,omitempty"`
	Total      int `json:"total,omitempty"`
	TotalPages int `json:"total_pages,omitempty"`
}

func OK(ctx *gin.Context, data interface{}, statusCode int) {
	ctx.JSON(statusCode, Response{
		Success: true,
		Data:    data,
	})
}

func Fail(ctx *gin.Context, status int, message string) {
	ctx.JSON(status, Response{
		Success: false,
		Error:   &ErrorInfo{Message: message},
	})
}
