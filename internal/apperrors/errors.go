package apperrors

import "errors"

var (
	ErrAlreadyExists = errors.New("record already exists")
)
