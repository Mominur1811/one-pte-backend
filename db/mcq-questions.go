package db

import (
	"encoding/json"
	sq "github.com/Masterminds/squirrel"
)

type McqOption struct {
	Id     int    `json:"id"`
	Option string `json:"option"`
}

type McqQuestion struct {
	Id            int             `json:"id"`
	Options       json.RawMessage `json:"options"` // [] McqOption
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

func (repo *McqQuestionRepo) GetMcqQuestionById(id int) (*McqQuestion, error) {
	query := GetQueryBuilder().
		Select("id", "options", "correctOption").
		From(repo.Table).
		Where(sq.Eq{"id": id})

	var result McqQuestion

	err := query.QueryRow().Scan(&result.Id, &result.Options, &result.CorrectOption)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
