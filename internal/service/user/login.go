package user

import (
	"context"
	"fmt"

	"rovioletta/todo-list-api/internal/apperrors"
	pwdcrypto "rovioletta/todo-list-api/internal/pkg/pwd-crypto"
)

func (s *Service) Login(ctx context.Context, login, password string) (token string, err error) {
	user, err := s.queries.GetUserByLogin(ctx, login)
	if err != nil {
		return "", fmt.Errorf("get user from db: %w", err)
	}

	equal, err := pwdcrypto.VerifyPassword(password, user.PasswordHash)
	if err != nil {
		return "", fmt.Errorf("failed to verify password: %w", err)
	}

	if equal == false {
		return "", apperrors.ErrWrongPassword
	}

	token, err = s.auth.Generate(login)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, err
}
