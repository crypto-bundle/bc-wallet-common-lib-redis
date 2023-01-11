package crypter

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/crypto-bundle/bc-wallet-common/pkg/env"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Key        []byte `envconfig:"RSA_ENCRYPTION_KEY"`
	KeyPath    string `envconfig:"RSA_ENCRYPTION_KEY_PATH" default:"./id_rsa"`
	NoEnvClean bool   `envconfig:"RSA_ENCRYPTION_NO_ENV_CLEAN"`

	privateKey *rsa.PrivateKey
}

//nolint:funlen // its ok
func (c *Config) Prepare() error {
	err := envconfig.Process("", c)
	if err != nil {
		return err
	}

	if !c.NoEnvClean {
		err = env.CleanByEnvTags("", c)
		if err != nil {
			return err
		}
	}

	key := c.Key
	if len(key) == 0 {
		key, err = os.ReadFile(c.KeyPath)
		if err != nil {
			return err
		}
	}

	keyBlock, _ := pem.Decode(key)
	if keyBlock == nil {
		return ErrInvalidKey
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if err != nil {
		return fmt.Errorf("%w: %q", ErrParsePK, err.Error())
	}
	c.privateKey = privateKey

	return nil
}

func (c *Config) GetPrivateKey() *rsa.PrivateKey {
	return c.privateKey
}
