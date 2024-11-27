package db

import (
	"database/sql"
)

var db *sql.DB

func GetDB() *sql.DB {
	return db
}

func StartTransaction() (*sql.Tx, error) {
	tx, err := GetDB().Begin()
	if err != nil {
		return nil, err
	}
	return tx, nil
}
