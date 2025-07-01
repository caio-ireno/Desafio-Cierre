package apperrors

import "errors"

var (
	ErrLoadCSV   = errors.New("failed to Load CSV")
	ErrEmptyData = errors.New("empty data")
	ErrNotFound  = errors.New("ticket not found in database")
)
