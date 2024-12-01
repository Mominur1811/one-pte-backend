package db

import (
	"encoding/json"
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
