package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
	"os"
)

// encryption
func Encrypt(data []byte) []byte {

	//get aes key from .env file
	key64 := os.Getenv("AES_KEY")

	//aes key is stored as base64 encoded,so decode that to get []byte data
	//Note : key is 32 bytes
	key, err := base64.StdEncoding.DecodeString(key64)
	if err != nil {
		log.Fatalf("Failed to decode AES_KEY: %v", err)
	}

	//aes.NewCipher() takes key as argument and returns a cipher block of size 16 bytes (128 bits).
	//we use 16, 24, or 32 bytes (i.e key size) to select AES-128, AES-192, or AES-256.
	//no.of blocks depends on the user input i.e ceil(input size / 16)

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("Failed to create AES cipher: %v", err)
	}

	//cipher.NewGCM() takes block as input and generates an authentication tag (16 bytes)
	// which is useful to prevent file tempering (content change or modification)
	// If the file is tampered, even by one byte, GCM decryption will detect it and reject the data.
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("Failed to create GCM cipher mode: %v", err)
	}

	//nonce is a rand number used once to make sure aec-gcm encryption is unique and secure
	//each and every block of single user input must use the same nonce number
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalf("Failed to generate nonce: %v", err)
	}

	//encrypts the data	and adds the result with nonce as first
	//eg :
	// Encrypted data (hex): 3b5569e6cfd2a6dce0101bd2a117c257355f7c5b55ce448b55d9701c4b1b3b8c4a
	// Nonce (hex):          3b5569e6cfd2a6dce0101bd2
	// Cipher only (hex):    a117c257355f7c5b55ce448b55d9701c4b1b3b8c4a

	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

// decryption
func Decrypt(enc []byte) []byte {
	// decoded, err := hex.DecodeString(enc)
	// if err != nil {
	// 	log.Printf("Hex decode error: %v", err)
	// 	return nil
	// }

	key64 := os.Getenv("AES_KEY")
	key, err := base64.StdEncoding.DecodeString(key64)
	if err != nil {
		log.Printf("Base64 decode error: %v", err)
		return nil
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Printf("AES cipher init error: %v", err)
		return nil
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Printf("GCM init error: %v", err)
		return nil
	}

	//checks if the encrypted data is less then the nonce size, if yes then the enc data is invalid or tampered data
	nonceSize := gcm.NonceSize()
	if len(enc) < nonceSize {
		log.Printf("Invalid encrypted data: too short")
		return nil
	}

	//since the encrypted first contains the nonce and then the actual encrypted data (user input)
	//so we need to split that into separate
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//decrypt the encrypted data (ciphertext) using nonce
	//if the data has been tampered with, decryption will fail and return an error.
	dec, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Printf("Decryption failed: %v", err)
		return nil
	}

	return dec
}
