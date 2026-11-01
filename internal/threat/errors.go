package threat

import "errors"

var (
	ErrNotFound      = errors.New("threat not found")
	ErrConflict      = errors.New("threat already exists")
	ErrInvalidMD5    = errors.New("invalid md5 format")
	ErrInvalidFields = errors.New("invalid fileds")
)
