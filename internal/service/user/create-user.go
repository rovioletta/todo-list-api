package user

import (
	"context"
	"fmt"

	pwdcrypto "rovioletta/todo-list-api/internal/pkg/pwd-crypto"
)

func (s *Service) CreateUser(ctx context.Context, login, password string) (userID uint64, err error) {
	hash, err := pwdcrypto.HashPassword(password)
	if err != nil {
		return userID, fmt.Errorf("hash password: %w", err)
	}

	userID, err = s.db.CreateUser(ctx, login, hash)
	if err != nil {
		return userID, fmt.Errorf("insert user to db: %w", err)
	}

	return
}