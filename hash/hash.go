package hash

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"log"
)

func CalculateHash(code string) []byte {
	hash := sha256.Sum256([]byte(code))
	return hash[:]
}

func Compare(code string, hash []byte) bool {
	return bytes.Equal(hash, CalculateHash(code))
}

func HexToByteArray(hexString string) []byte {
	// Remove "0x" prefix if present
	if len(hexString) >= 2 && hexString[:2] == "0x" {
		hexString = hexString[2:]
	}

	// Decode hex string to byte slice
	byteArray, err := hex.DecodeString(hexString)
	if err != nil {
		log.Fatalf("failed to decode hex string: %v", err)
	}

	return byteArray
}
