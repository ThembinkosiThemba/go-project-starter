package http

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func Context() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx, cancel
}

type Response struct {
	Status  int         `json:"status"`
	AppCode int         `json:"appCode"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func WriteJSON(c *gin.Context, status int, appCode int, data interface{}, message string) {
	c.Header("Content-Type", "application/json")
	c.JSON(status, Response{Status: status, AppCode: appCode, Message: message, Data: data})
}
