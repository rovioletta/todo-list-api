package db

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	*Queries
	logger *slog.Logger
	dbpool *pgxpool.Pool
}

func NewDB(logger *slog.Logger) *DB {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Error("Unable to connect to database", slog.String("error", err.Error()))
		os.Exit(1)
	}

	return &DB{
		Queries: New(dbpool),
		logger:  logger,
		dbpool:  dbpool,
	}
}

func (db *DB) CloseDB() {
	db.logger.Info("Closing database connection...")
	db.dbpool.Close()
}
