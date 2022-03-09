package logger

import "errors"

var (
	ErrNamedLoggerAlreadyRegistered = errors.New("logger with passed name already registered")
)

