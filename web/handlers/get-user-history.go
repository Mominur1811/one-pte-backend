package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"one-pte-backend/db"
	"one-pte-backend/web/utils"
	"strconv"
)

func GetUserHistory(g *gin.Context) {
	useId, err := strconv.Atoi(g.Param("userId"))
	if err != nil {
		utils.SendError(g, http.StatusBadRequest, "Failed to get question details", "Invalid question ID")
		return
	}
	qType := g.DefaultQuery("qType", "")

	history, err := db.GetUserHistoryRepo().GetUserHistory(useId, qType)
	if err != nil {
		utils.SendError(g, http.StatusInternalServerError, "Failed to get user history", "")
		return
	}

	utils.SendData(g, history)
}
