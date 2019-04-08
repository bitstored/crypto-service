package service

import (
	"context"
	"fmt"
	"github.com/bitstored/crypto-service/pkg/crypto/pbkdf3"
)

const (
	minIter = 4096
	saltLen = 32
)

type CryptoService struct{}

func (cs *CryptoService) EncryptPassword(ctx context.Context, password []byte, salt []byte, iter int) (hash []byte, err error) {
	if iter < minIter {
		return nil, fmt.Errorf("minimum number of iterations is %d", minIter)
	}
	hash, err = pbkdf3.HashWithSHA512(password, salt, iter, len(password))
	return hash, err
}

func NewCryptoService() *CryptoService {
	return &CryptoService{}
}
