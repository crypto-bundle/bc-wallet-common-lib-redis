package crypto

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

type Config struct {
	KeyPath string `envconfig:"key_path"  default:"./id_rsa"`
}

//nolint:funlen // its ok
func (c *Config) Prepare() error {
	_, err := os.Stat(c.KeyPath)
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
		return nil, fmt.Errorf("fail to get idrsa, invalid key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("fail to get idrsa, %s", err.Error())
	}

	return privateKey, nil
}
