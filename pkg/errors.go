package pkg

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrNotFound            = errors.New("resource not found")
	ErrBadParamInput       = errors.New("param is not valid")
	ErrAuthFailed          = errors.New("authorization failed")
)
