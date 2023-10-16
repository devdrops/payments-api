package database

import (
	"errors"
)

var (
	ErrNotExist     = errors.New("Row does not exist.")
	ErrDuplicate    = errors.New("Row already exists.")
	ErrReadFailed   = errors.New("Failed to read row.")
	ErrCreateFailed = errors.New("Failed to create row.")
)
