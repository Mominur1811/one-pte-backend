package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"one-pte-backend/db"
	"one-pte-backend/web/utils"
	"strconv"
)

func GetQuestionDetails(g *gin.Context) {

	questionId, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		utils.SendError(g, http.StatusBadRequest, "Failed to get question details", "Invalid question ID")
		return
	}

	qType := g.DefaultQuery("type", "")
	if qType == "" {
		utils.SendError(g, http.StatusBadRequest, "Failed get question details", "Question type is required")
		return
	}

	title, err := db.GetQuestionRepo().GetTitleById(questionId)
	if err != nil {
		utils.SendError(g, http.StatusInternalServerError, "Failed to get question details", "")
		return
	}

	var result interface{}

	switch qType {
	case "SST":
		result, err = getSSTQuestionDetails(questionId, title)
	case "RO":
		result, err = getROQuestionDetails(questionId, title)
	case "MCQ":
		result, err = getMCQQuestionDetails(questionId, title)
	default:
		utils.SendError(g, http.StatusBadRequest, "Failed get question details", "Invalid question type")
		return
	}

	if err != nil {
		utils.SendError(g, http.StatusInternalServerError, "Failed to get question details", "")
		return
	}

	utils.SendData(g, result)
}

func getSSTQuestionDetails(questionId int, title string) (interface{}, error) {

	sstQues, err := db.GetSSTQuestionsRepo().GetSSTQuestionDetails(questionId)
	if err != nil {
		return nil, err
	}

	var audioInfos []db.AudioInfo

	err = json.Unmarshal(sstQues.Audio, &audioInfos)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id":                questionId,
		"title":             title,
		"questionTimeLimit": sstQues.QuestionTimeLimit,
		"audios":            audioInfos,
	}, nil
}

func getROQuestionDetails(questionId int, title string) (interface{}, error) {
	roQues, err := db.GetRoQuestionsRepo().GetRoQuestionById(questionId)
	if err != nil {
		return nil, err
	}

	var paragraphs []db.Paragraph
	if err := json.Unmarshal(roQues.Paragraph, &paragraphs); err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id":           questionId,
		"title":        title,
		"paragraph":    paragraphs,
		"correctOrder": roQues.CorrectOrder,
	}, nil

}

func getMCQQuestionDetails(questionId int, title string) (interface{}, error) {
	mcqQues, err := db.GetMcqQuestionRepo().GetMcqQuestionById(questionId)
	if err != nil {
		return nil, err
	}

	var options []db.McqOption
	if err := json.Unmarshal(mcqQues.Options, &options); err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id":            questionId,
		"title":         title,
		"options":       options,
		"correctOption": mcqQues.CorrectOption,
	}, nil
}
