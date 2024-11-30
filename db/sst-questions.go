package db

import (
	"database/sql"
	"encoding/json"
	"github.com/guregu/null"
)

type AudioInfo struct {
	Speaker string `json:"speaker"`
	Path    string `json:"path"`
}

type SSTQuestions struct {
	Id                null.Int        `json:"id"`
	QuestionTimeLimit int             `json:"questionTimeLimit"`
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
