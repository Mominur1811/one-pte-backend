package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "Success",
		"data":    data,
	})
}
