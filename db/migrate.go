package db

import (
	"log/slog"
	"one-pte-backend/config"

	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

func MigrateDB() {
	conf := config.GetConfig()

	// connect to database first
	ConnectDB()

	// try migrating tables
	migrations := &migrate.FileMigrationSource{
		Dir: conf.MigrationSource,
	}

	_, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		panic(err)
	}

	slog.Info("Successfully migrated database")
}
