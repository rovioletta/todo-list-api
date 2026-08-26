package db

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"rovioletta/todo-list-api/internal/apperrors"
)

type queriesInterface interface {
	CreateUser(ctx context.Context, login string, passwordHash string) (uint64, error)
	GetUserByLogin(ctx context.Context, login string) (GetUserByLoginRow, error)
	CreateTask(ctx context.Context, arg *CreateTaskParams) (uint64, error)
	WithTx(tx pgx.Tx) *Queries
}

type errorHandler struct {
	queries queriesInterface
}

func NewErrorHandler(queries queriesInterface) *errorHandler {
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

	return err
}
