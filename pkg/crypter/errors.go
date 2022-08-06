package crypter

import (
	"errors"
)

var (
	ErrInvalidKey = errors.New("fail to get idrsa, invalid key")
	ErrParsePK    = errors.New("fail to parse private key")
)
