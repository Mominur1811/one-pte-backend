package handlers

import (
	"github.com/gin-gonic/gin"
	"one-pte-backend/db"
)

func GetUserHistory(g *gin.Context) {
	useId := g.Param("userId")
	qType := g.DefaultQuery("qType", "")

	history, err := db.GetUserHistoryRepo().GetUserHistory(useId, qType)
}
