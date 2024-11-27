package utils

import "github.com/gin-gonic/gin"

func SendError(ctx *gin.Context, status int, message string, data interface{}) {
	ctx.JSON(status, map[string]interface{}{
		"status":  false,
		"message": message,
		"data":    data,
	})
}
