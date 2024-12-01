package handlers

import (
	"github.com/gin-gonic/gin"
	"one-pte-backend/web/utils"
)

func Hello(c *gin.Context) {
	utils.SendData(c, "OK")
}
