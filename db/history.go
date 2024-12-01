package db

import (
	"encoding/json"
	sq "github.com/Masterminds/squirrel"
	"time"
)

type UserHistory struct {
	ID           int             `json:"id"`
	UserID       int             `json:"userId"`
	QuestionID   int             `json:"questionId"`
	QuestionType string          `json:"questionType"`
	Answer       json.RawMessage `json:"answer"`
	ObtainMarks  float64         `json:"obtainMarks"`
	TotalMarks   float64         `json:"totalMarks"`
	CreatedAt    time.Time       `json:"createdAt"`
}

type UserHistoryRepo struct {
	Table string
}

var userHistoryRepo *UserHistoryRepo

func InitUserHistoryRepo() {
	userHistoryRepo = &UserHistoryRepo{
		Table: "user_history",
	}
}

func GetUserHistoryRepo() *UserHistoryRepo {
	return userHistoryRepo
}

func (repo *UserHistoryRepo) InsertUserHistory(userHistory UserHistory) (*UserHistory, error) {
	query := GetQueryBuilder().
		Insert(repo.Table).
		Columns("user_id", "question_id", "question_type",
			"answer", "obtain_marks",
			"total_marks", "created_at").
		Values(
			userHistory.UserID,
			userHistory.QuestionID,
			userHistory.QuestionType,
			userHistory.Answer,
			userHistory.ObtainMarks,
			userHistory.TotalMarks,
			userHistory.CreatedAt,
		).
		Suffix("RETURNING *")

	var result UserHistory

	err := query.QueryRow().Scan(
		&result.ID,
		&result.UserID,
		&result.QuestionID,
		&result.QuestionType,
		&result.Answer,
		&result.ObtainMarks,
		&result.TotalMarks,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (repo *UserHistoryRepo) GetUserHistory(userId int, qType string) ([]UserHistory, error) {
	query := GetQueryBuilder().
		Select("").
		From(repo.Table).
		Where(sq.Eq{"user_id": userId})
	if qType != "" {
		query = query.Where(sq.Eq{"question_type": qType})
	}

	rows, err := query.Query()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []UserHistory

	for rows.Next() {
		var row UserHistory
		err := rows.Scan(row.ID, row.UserID, row.QuestionID, row.QuestionType, row.Answer, row.ObtainMarks, row.TotalMarks, row.CreatedAt)
		if err != nil {
			return nil, err
		}

		result = append(result, row)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
