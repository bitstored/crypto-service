package service

import (
	"context"
	"fmt"
	"github.com/bitstored/crypto-service/pkg/crypto/pbkdf3"
	"github.com/bitstored/crypto-service/pkg/crypto/twofish"
)

const (
	minIter = 4096
	saltLen = 32
)

type CryptoService struct{}

func (cs *CryptoService) EncryptFile(ctx context.Context, content, secretPhrase []byte) ([]byte, error) {
	c, err := twofish.NewCipher(secretPhrase)
	if err != nil {
		return nil, err
	}

	out := make([]byte, len(content))
	c.Encrypt(out, content)

	return out, nil
}

func (cs *CryptoService) DecryptFile(ctx context.Context, content, secretPhrase []byte) ([]byte, error) {
	c, err := twofish.NewCipher(secretPhrase)
	if err != nil {
		return nil, err
	}

	out := make([]byte, len(content))
	c.Decrypt(out, content)

	return out, nil
}

func (cs *CryptoService) EncryptPassword(ctx context.Context, password []byte, salt []byte, iter int) (hash []byte, err error) {
	if iter < minIter {
		return nil, fmt.Errorf("minimum number of iterations is %d", minIter)
	}
	hash = pbkdf3.HashWithSHA512(password, salt, iter, len(password))
	return hash, err
}

func (cs *CryptoService) Encrypt(ctx context.Context, data, salt []byte) ([]byte, error) {
	c, err := twofish.NewCipher(salt)
	if err != nil {
		return nil, err
	}

	out := make([]byte, len(data))
	c.Encrypt(out, data)

	return out, nil
}

func (cs *CryptoService) Decrypt(ctx context.Context, data, salt []byte) ([]byte, error) {
	c, err := twofish.NewCipher(salt)
	if err != nil {
		return nil, err
	}

	out := make([]byte, len(data))
	c.Decrypt(out, data)

	return out, nil
}
func NewCryptoService() *CryptoService {
	return &CryptoService{}
}
