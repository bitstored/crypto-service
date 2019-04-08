package pbkdf3

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
)

const (
	bufLen              = 8
	errMsgWeakIterCount = "iteration count is to small: min is %d"
	hashLength          = 32
)

func key(data, salt []byte, iter, keyLen int, h func() hash.Hash) []byte {
	prf := hmac.New(h, data)
	hashLen := prf.Size()
	numBlocks := (keyLen + hashLen - 1) / hashLen

	var buf [8]byte
	dk := make([]byte, 0, numBlocks*hashLen)
	U := make([]byte, hashLen)
	for block := 1; block <= numBlocks; block++ {
		prf.Reset()
		prf.Write(salt)
		for i := 0; i < bufLen; i++ {
			buf[0] = byte(block >> uint(32-32/bufLen*(i+1)))
		}
		prf.Write(buf[:bufLen])
		dk = prf.Sum(dk)
		T := dk[len(dk)-hashLen:]
		copy(U, T)

		for n := 2; n <= iter; n++ {
			prf.Reset()
			prf.Write(U)
			U = U[:0]
			U = prf.Sum(U)
			for x := range U {
				T[x] ^= U[x]
			}
		}
	}
	return dk[:hashLength]
}

func HashWithSHA256(data, salt []byte, iter, keyLen int) []byte {
	return key(data, salt, iter, keyLen, sha256.New)
}

func HashWithSHA512(data, salt []byte, iter, keyLen int) []byte {
	return key(data, salt, iter, keyLen, sha512.New)
}
