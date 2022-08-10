package crypter

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
)

type Service struct {
	cfg config
}

func New(cfg config) *Service {
	return &Service{cfg: cfg}
}

func (s *Service) Encrypt(msg string) (string, error) {
	privateKey := s.cfg.GetPrivateKey()
	encMsgBytes, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &privateKey.PublicKey, []byte(msg), nil)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(encMsgBytes), nil
}

func (s *Service) Decrypt(encMsg string) ([]byte, error) {
	encMsgBytes, err := hex.DecodeString(encMsg)
	if err != nil {
		return nil, err
	}

	msgBytes, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, s.cfg.GetPrivateKey(), encMsgBytes, nil)
	if err != nil {
		return nil, err
	}

	return msgBytes, nil
}
