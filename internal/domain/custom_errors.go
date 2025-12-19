package domain

import "errors"

var (
	ErrInvalidId        = errors.New("id must be greater than zero")
	ErrInvalidMonth     = errors.New("monthId must be greater than 0 and less than 13")
	ErrInvalidAmount    = errors.New("amount must be greater than zero")
	ErrEmptyDescription = errors.New("empty description")
	ErrEmptyCategory    = errors.New("empty category")
)
