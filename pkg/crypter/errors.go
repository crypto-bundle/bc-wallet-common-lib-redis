package crypter

import (
	"fmt"
)

var (
	ErrInvalidKey = fmt.Errorf("fail to get idrsa, invalid key")
	ErrParsePK    = fmt.Errorf("fail to parse private key")
)
