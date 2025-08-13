package usecase

import "errors"

var (
	ErrNotFound = errors.New("task not found")
	ErrInvalid  = errors.New("invalid input")
)
