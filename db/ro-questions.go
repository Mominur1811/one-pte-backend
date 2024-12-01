package db

import (
	"encoding/json"
	sq "github.com/Masterminds/squirrel"
)

type Paragraph struct {
	Id   int    `json:"id" validate:"required,gte=0"`
	Test string `json:"test" validate:"required"`
}

type RoQuestions struct {
	Id           int             `json:"id" validate:"required"`
	Paragraph    json.RawMessage `json:"paragraph" validate:"required"` // []Paragraph
	CorrectOrder []int           `json:"correctOrder" validate:"required"`
}

type RoQuestionRepo struct {
	Table string
}

var roQuestionRepo *RoQuestionRepo

func InitRoQuestionRepo() {
	roQuestionRepo = &RoQuestionRepo{
		Table: "ro_questions",
	}
}

func GetRoQuestionsRepo() *RoQuestionRepo {
	return roQuestionRepo
}

func (repo *RoQuestionRepo) InsertRoQuestion(quest *RoQuestions) (*RoQuestions, error) {

	query := GetQueryBuilder().
		Insert(repo.Table).
		Columns("id", "paragraph", "correct_order").
		Values(quest.Id, quest.Paragraph, quest.CorrectOrder).
		Suffix("RETURNING *")

	var result RoQuestions

	err := query.QueryRow().
		Scan(&result.Id, &result.Paragraph, &result.CorrectOrder)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (repo *RoQuestionRepo) GetRoQuestionById(id int) (*RoQuestions, error) {
	query := GetQueryBuilder().
		Select("id", "paragraph", "correct_order").
		From(repo.Table).
		Where(sq.Eq{"id": id})

	var result RoQuestions

	err := query.QueryRow().Scan(&result.Id, &result.Paragraph, &result.CorrectOrder)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
