package domain

import "errors"

var (
	ErrAccountNotFound = errors.New("account not found")
	ErrDuplicatedKey   = errors.New("api key already exists")
	ErrInvoiceNotFound = errors.New("invoice not found")
	ErrUnauthorized    = errors.New("unauthorized access")

	ErrInvalidAmount = errors.New("invalid amount")
	ErrInvalidStatus = errors.New("invalid status")
)
