package crypter

import (
	"crypto/rsa"
)

type config interface {
	GetKey() (*rsa.PrivateKey, error)
}
