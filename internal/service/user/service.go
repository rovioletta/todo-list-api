package user

import (
	"context"

	"rovioletta/todo-list-api/internal/db"
	"rovioletta/todo-list-api/internal/pkg/tokens"
)

type queries interface {
	CreateUser(ctx context.Context, login string, passwordHash string) (userID uint64, err error)
	GetUserByLogin(ctx context.Context, login string) (user db.GetUserByLoginRow, err error)
}

type auth interface {
	Generate(login string) (token string, err error)
	Verify(accessToken string) (*tokens.CustomClaims, error)
}

type Service struct {
	queries queries
	auth    auth
}

func NewService(queries queries, auth auth) *Service {
	return &Service{
		queries: queries,
		auth:    auth,
	}
}
