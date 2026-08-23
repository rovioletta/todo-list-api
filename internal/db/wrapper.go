package db

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"rovioletta/todo-list-api/internal/apperrors"
)

type errorHandler struct {
	db DBTX
}

func NewErrorHandler(db DBTX) *errorHandler {
	return &errorHandler{db: db}
}

func (w *errorHandler) Exec(ctx context.Context, query string, params ...interface{}) (pgconn.CommandTag, error) {
	res, err := w.db.Exec(ctx, query, params...)
	return res, handleError(err)
}

func (w *errorHandler) Query(ctx context.Context, query string, params ...interface{}) (pgx.Rows, error) {
	res, err := w.db.Query(ctx, query, params...)
	return res, handleError(err)
}

func (w *errorHandler) QueryRow(ctx context.Context, query string, params ...interface{}) pgx.Row {
	return w.db.QueryRow(ctx, query, params...)
}

func handleError(err error) error {
	if err == nil {
		return nil
	}

	if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
		if pgErr.Code == "23505" { // unique_violation
			return apperrors.ErrAlreadyExists
		}
	}

	return err
}
