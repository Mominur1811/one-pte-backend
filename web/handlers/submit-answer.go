package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"one-pte-backend/db"
	"one-pte-backend/web/utils"
	"time"
)

type Answer struct {
	UserID       int         `json:"userId"`
	QuestionID   int         `json:"questionId"`
	QuestionType string      `json:"questionType"`
	Answer       interface{} `json:"answer"`
}

func SubmitAnswer(g *gin.Context) {
	var answer Answer
	if err := g.ShouldBindJSON(&answer); err != nil {
		utils.SendError(g, http.StatusBadRequest, "Failed to submit answer", "")
		return
	}

	userHistory, err := processAnswer(answer)
	if err != nil {
		utils.SendError(g, http.StatusBadRequest, "Failed to submit answer", "")
		return
	}

	userHistory, err = db.GetUserHistoryRepo().InsertUserHistory(userHistory)
	if err != nil {
		utils.SendError(g, http.StatusBadRequest, "Failed to submit answer", "")
		return
	}

	utils.SendData(g, userHistory)

}

func processAnswer(answer Answer) (*db.UserHistory, error) {
	var err error
	var answerJSON []byte
	var obtainMarks, totalMarks float64

	switch answer.QuestionType {
	case "SST":
		obtainMarks, totalMarks = processSSTAnswer(answer)
		answerJSON, err = json.Marshal(answer.Answer)
	case "RO":
		obtainMarks, totalMarks, err = processROAnswer(answer)
		answerJSON, err = json.Marshal(answer.Answer)
	case "MCQ":
		obtainMarks, totalMarks, err = processMCQAnswer(answer)
		answerJSON, err = json.Marshal(answer.Answer)
	default:
		return nil, fmt.Errorf("unsupported question type")
	}

	if err != nil {
		return nil, err
	}

	userHistory := &db.UserHistory{
		UserID:       answer.UserID,
		QuestionID:   answer.QuestionID,
		QuestionType: answer.QuestionType,
		Answer:       answerJSON,
		ObtainMarks:  obtainMarks,
		TotalMarks:   totalMarks,
		CreatedAt:    time.Now(),
	}

	return userHistory, nil
}

func processSSTAnswer(answer Answer) (float64, float64) {
	rand.Seed(time.Now().UnixNano())
	marks := 1 + rand.Float64()*(10-1)
	totalMarks := 10.0

	return marks, totalMarks
}

func processROAnswer(answer Answer) (float64, float64, error) {

	roQues, err := db.GetRoQuestionsRepo().GetRoQuestionById(answer.QuestionID)
	if err != nil {
		return 0, 0, err
	}

	totalMarks := float64(len(roQues.CorrectOrder) - 1)

	submittedOrder, ok := utils.ConvertToIntSlice(answer.Answer)
	if !ok {
		return 0, totalMarks, fmt.Errorf("invalid answer format")
	}

	obtainMarks := utils.CalculateCorrectAdjacentPairs(roQues.CorrectOrder, submittedOrder)

	return obtainMarks, totalMarks, nil
}

func processMCQAnswer(answer Answer) (float64, float64, error) {
	mcqQues, err := db.GetMcqQuestionRepo().GetMcqQuestionById(answer.QuestionID)
	if err != nil {
		return 0, 0, err
	}

	totalMarks := 1.0

	selectedOption, ok := answer.Answer.(float64)
	if !ok {
		return 0, 0, fmt.Errorf("invalid MCQ answer format")
	}

	obtainsMarks := -1.0

	if int(selectedOption) == mcqQues.CorrectOption {
		obtainsMarks = 1.0
	}

	return obtainsMarks, totalMarks, nil
}
