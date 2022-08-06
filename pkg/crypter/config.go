package crypter

import (
	"fmt"
	"os"

	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/kelseyhightower/envconfig"
)

const ConfigPrefix = "RSA"

type Config struct {
	KeyPath string `envconfig:"RSA_ENCRYPTION_KEY_PATH"  default:"./id_rsa"`
}

//nolint:funlen // its ok
func (c *Config) Prepare() error {
	err := envconfig.Process(ConfigPrefix, c)
	if err != nil {
		return err
	}

	_, err = os.Stat(c.KeyPath)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) GetKey() (*rsa.PrivateKey, error) {
	key, err := os.ReadFile(c.KeyPath)
	if err != nil {
		return nil, err
	}

	keyBlock, _ := pem.Decode(key)
	if keyBlock == nil {
		return nil, ErrInvalidKey
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("%w: %q", ErrParsePK, err.Error())
	}

	return privateKey, nil
}
