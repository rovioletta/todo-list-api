package apperrors

import "errors"

var (
	ErrAlreadyExists = errors.New("record already exists")
	ErrWrongPassword = errors.New("wrong password")
	ErrNotFound      = errors.New("not found")
)
