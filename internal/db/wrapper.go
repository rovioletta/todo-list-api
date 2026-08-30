package db

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"rovioletta/todo-list-api/internal/apperrors"
)

type errorHandler struct {
	queries *Queries
}

func NewErrorHandler(queries *Queries) *errorHandler {
	return &errorHandler{queries: queries}
}

func (w *errorHandler) CreateUser(ctx context.Context, login string, passwordHash string) (uint64, error) {
	res, err := w.queries.CreateUser(ctx, login, passwordHash)
	return res, handleError(err)
}

func (w *errorHandler) GetUserByLogin(ctx context.Context, login string) (GetUserByLoginRow, error) {
	res, err := w.queries.GetUserByLogin(ctx, login)
	return res, handleError(err)
}

func (w *errorHandler) CreateTask(ctx context.Context, arg *CreateTaskParams) (uint64, error) {
	res, err := w.queries.CreateTask(ctx, arg)
	return res, handleError(err)
}

func (w *errorHandler) UpdateTask(ctx context.Context, arg *UpdateTaskParams) (uint64, error) {
	res, err := w.queries.UpdateTask(ctx, arg)
	return res, handleError(err)
}

func (w *errorHandler) GetTaskByID(ctx context.Context, id uint64) (Task, error) {
	res, err := w.queries.GetTaskByID(ctx, id)
	return res, handleError(err)
}

func (w *errorHandler) WithTx(tx pgx.Tx) *Queries {
	return w.queries.WithTx(tx)
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

	if err == pgx.ErrNoRows {
		return apperrors.ErrNotFound
	}

	return err
}
