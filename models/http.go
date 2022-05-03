package models

import "errors"

// APIResponse ...
type APIResponse struct {
	Status    string      `json:"status"`
	Message   string      `json:"message"`
	ErrorCode string      `json:"error_code"`
	Data      interface{} `json:"data"`
}

var (
	ErrValidationFailed    = errors.New("request input body validation failed")
	ErrNoDataAvailable     = errors.New("no data available for requested input")
	ErrPasswordReset       = errors.New("reset password")
	ErrInvalidSession      = errors.New("invalid session")
	ErrInternalServerError = errors.New("internal server error")
)
