package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"os"
)

func Encrypt(data []byte) string {
	key64 := os.Getenv("AES_KEY")
	key, _ := base64.StdEncoding.DecodeString(key64)
	block, _ := aes.NewCipher([]byte(key))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	ciphertxt := gcm.Seal(nonce, nonce, data, nil)
	enc := hex.EncodeToString(ciphertxt)
	return enc
}
