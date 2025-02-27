package storage

import "errors"

var (
	ErrURLNotFound = errors.New("URL not found")
	ErrUrlExists   = errors.New("URL already exists")
)
