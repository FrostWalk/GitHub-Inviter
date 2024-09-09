package hash

import (
	"bytes"
	"crypto/sha256"
)

func CalculateHash(password string) []byte {
	h := sha256.New()
	h.Write([]byte(password))
	return h.Sum(nil)
}

func Compare(password string, hash []byte) bool {
	return bytes.Equal(hash, CalculateHash(password))
}
