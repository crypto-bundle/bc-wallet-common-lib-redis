package crypter

import (
	"crypto/rsa"
)

type config interface {
	GetPrivateKey() *rsa.PrivateKey
}
