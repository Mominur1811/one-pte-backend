package handlers

import (
	"github.com/gin-gonic/gin"
	"one-pte-backend/db"
	"one-pte-backend/web/utils"
)

func AddSstQuestion(g *gin.Context) {
	var sstQues db.SSTQuestions

	if err := g.ShouldBindJSON(&sstQues); err != nil {
		utils.SendError(g, 400, "Failed to add question", "")
		return
	}

	if err := utils.ValidateStruct(sstQues); err != nil {
		utils.SendError(g, 400, "Failed to add question", "")
		return
	}

	tx, err := db.StartTransaction()
	if err != nil {
		utils.SendError(g, 500, "Failed to add question", "")
		return
	}

	sstQues.Type = "SST"

	qId, err := db.GetQuestionRepo().CreateQuestionTx(tx, sstQues.Type, sstQues.Title)
	if err != nil {
		utils.SendError(g, 500, "Failed to add question", "")
		return
	}

	sstQues.Id = *qId

	err = db.GetSSTQuestionsRepo().InsertSSTQuestions(tx, sstQues)
	if err != nil {
		utils.SendError(g, 500, "Failed to add question", "")
		return
	}

	err = tx.Commit()
	if err != nil {
		utils.SendError(g, 500, "Failed to add question", "")
		return
	}

}
