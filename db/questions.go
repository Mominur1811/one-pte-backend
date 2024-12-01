package db

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/guregu/null"
)

type Question struct {
	Id       int       `json:"id"`
	Title    string    `json:"title"`
	Type     string    `json:"type"`
	CreateAt null.Time `json:"create_at"`
}

type QuestionRepo struct {
	Table string
}

var questionRepo *QuestionRepo

func InitQuestionRepo() {
	questionRepo = &QuestionRepo{
		Table: "questions",
	}
}

func GetQuestionRepo() *QuestionRepo {
	return questionRepo
}

func (repo *QuestionRepo) CreateQuestionTx(tx *sql.Tx, qType, title string) (*null.Int, error) {
	var id null.Int

	query := GetQueryBuilder().
		Insert(repo.Table).
		Columns("type", "title").
		Values(qType, qType).
		Suffix("RETURNING id")

	err := query.RunWith(tx).QueryRow().Scan(&id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (repo *QuestionRepo) GetQuestionsList(qType string) ([]Question, error) {

	query := GetQueryBuilder().
		Select("*").
		From(repo.Table)
	if qType != "" {
		query = query.Where(sq.Eq{"type": qType})
	}

	var questions []Question

	rows, err := query.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var q Question

		err := rows.Scan(&q.Id, &q.Title, &q.Type, &q.CreateAt)
		if err != nil {
			return nil, err
		}

		questions = append(questions, q)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return questions, nil
}

func (repo *QuestionRepo) GetTitleById(id int) (string, error) {

	query := GetQueryBuilder().
		Select("title").
		From(repo.Table).
		Where(sq.Eq{"id": id})

	var title string

	err := query.QueryRow().Scan(&title)
	if err != nil {
		return "", err
	}

	return title, nil
}
