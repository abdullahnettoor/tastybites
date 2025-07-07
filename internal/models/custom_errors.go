package models

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")

	ErrUnauthorized = errors.New("unauthorized")

	ErrIsEmpty = errors.New("data is empty")
)
