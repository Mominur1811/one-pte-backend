package db

import (
	"database/sql"
	"encoding/json"
	sq "github.com/Masterminds/squirrel"
	"github.com/guregu/null"
)

type AudioInfo struct {
	Speaker string `json:"speaker"`
	Path    string `json:"path"`
}

type SSTQuestions struct {
	Id                null.Int        `json:"id"`
	QuestionTimeLimit int             `json:"questionTimeLimit"` //[]AudioInfo
	Audio             json.RawMessage `json:"audio"`
}

type SSTQuestionsRepo struct {
	Table string
}

var sstQuestionsRepo *SSTQuestionsRepo

func InitSSTQuestionsRepo() {
	sstQuestionsRepo = &SSTQuestionsRepo{
		Table: "sst_questions",
	}
}

func GetSSTQuestionsRepo() *SSTQuestionsRepo {
	return sstQuestionsRepo
}

func (repo *SSTQuestionsRepo) InsertSSTQuestions(tx *sql.Tx, sstQuestions SSTQuestions) error {
	query := GetQueryBuilder().
		Insert(repo.Table).
		Columns("id", "questionTimeLimit", "audio").
		Values(sstQuestions.Id, sstQuestions.QuestionTimeLimit, sstQuestions.Audio)

	err := query.RunWith(tx).QueryRow().Scan()
	if err != nil {
		return err
	}
	return nil
}

func (repo *SSTQuestionsRepo) GetSSTQuestionDetails(id int) (*SSTQuestions, error) {
	query := GetQueryBuilder().
		Select("id", "questionTimeLimit", "audio").
		From(repo.Table).
		Where(sq.Eq{"id": id})

	var sstQuestions SSTQuestions

	err := query.QueryRow().Scan(sstQuestions.Id, sstQuestions.QuestionTimeLimit, sstQuestions.Audio)
	if err != nil {
		return nil, err
	}

	return &sstQuestions, nil
}
