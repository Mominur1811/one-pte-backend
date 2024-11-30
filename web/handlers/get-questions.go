package handlers

import (
	"github.com/gin-gonic/gin"
	"one-pte-backend/db"
	"one-pte-backend/web/utils"
)

func GetQuestions(g *gin.Context) {

	qType := g.Param("type")

	questions, err := db.GetQuestionRepo().GetQuestions(qType)
	if err != nil {
		utils.SendError(g, 502, "Failed to get questions", "")
		return
	}

	utils.SendData(g, questions)
}
