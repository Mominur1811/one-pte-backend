package db

import (
	"encoding/json"
)

type McqOpstion struct {
	Id     int    `json:"id"`
	Option string `json:"option"`
}

type McqQuestion struct {
	Id            int             `json:"id"`
	Options       json.RawMessage `json:"options"`
	CorrectOption int             `json:"correctOption"`
}

type McqQuestionRepo struct {
	Table string
}

var mcqQuestionRepo *McqQuestionRepo

func InitMcqQuestionRepo() {
	mcqQuestionRepo = &McqQuestionRepo{
		Table: "mcq_questions",
	}
}

func GetMcqQuestionRepo() *McqQuestionRepo {
	return mcqQuestionRepo
}

func (repo *McqQuestionRepo) InsertMcqQuestion(question McqQuestion) (*McqQuestion, error) {

	query := GetQueryBuilder().
		Insert(repo.Table).
		Columns("id", "options", "correctOption").
		Values(question.Id, question.Options, question.CorrectOption).
		Suffix("RETURNING *")

	var result McqQuestion

	err := query.QueryRow().Scan(&result.Id, &result.Options, &result.CorrectOption)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
