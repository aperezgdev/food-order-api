package domain_errors

import "errors"

var (
	NotFound = errors.New("Not Found Element")
	Database = errors.New("Database error")
)
