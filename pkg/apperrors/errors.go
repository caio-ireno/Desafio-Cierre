package apperrors

import "errors"

var (
	ErrLoadCSV     = errors.New("failed to Load CSV")
	ErrEmptyData   = errors.New("empty data")
	ErrNotFound    = errors.New("ticket not found in database")
	ErrToConnectDB = errors.New("error to connect in database")
	ErrQueryDB     = errors.New("error querying database")
	ErrScanDB      = errors.New("error scanning database row")
)
