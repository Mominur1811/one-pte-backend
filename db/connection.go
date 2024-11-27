package db

import (
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"log/slog"
	"one-pte-backend/config"
	"os"
	"time"
)

func GetConnectionString(dbConf config.DB) string {
	connectionString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s",
		dbConf.User,
		dbConf.Pass,
		dbConf.Host,
		dbConf.Port,
		dbConf.Name,
	)
	if !dbConf.EnableSSLMode {
		connectionString += " sslmode=disable"
	}
	return connectionString
}

func connect(dbConf config.DB) *sql.DB {
	dbSource := GetConnectionString(dbConf)

	dbCon, err := sql.Open("postgres", dbSource)
	if err != nil {
		slog.Error(fmt.Sprintf("Connection error: %v", err))
		os.Exit(1)
	}
	if err := dbCon.Ping(); err != nil {
		slog.Error(fmt.Sprintf("Db Ping error: %v", err))
		os.Exit(1)
	}

	dbCon.SetConnMaxIdleTime(
		time.Duration(dbConf.MaxIdleTimeInMinute * int(time.Minute)),
	)

	return dbCon
}

func ConnectDB() {
	conf := config.GetConfig()

	db = connect(conf.Db)
	slog.Info("Connected to read database")

	psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}

func CloseDB() {
	if err := db.Close(); err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Info("Disconnected from database")

}
