package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

// AES-256 key (32 bytes)
var key = []byte("12345678901234567890123456789012")

// ========================================
// ENCRYPT FILE
// ========================================

func EncryptFile(inputPath string) error {

	// Read file
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return err
	}

	// Create AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Create GCM
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// Create nonce
	nonce := make([]byte, gcm.NonceSize())

	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return err
	}

	// Encrypt data
	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	// Output file
	output := inputPath + ".enc"

	// Save encrypted file
	err = os.WriteFile(output, ciphertext, 0644)
	if err != nil {
		return err
	}

	return nil
}

// ========================================
// DECRYPT FILE
// ========================================

func DecryptFile(inputPath string) error {

	// Read encrypted file
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return err
	}

	// Create AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Create GCM
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// Extract nonce
	nonceSize := gcm.NonceSize()

	nonce := data[:nonceSize]
	ciphertext := data[nonceSize:]

	// Decrypt
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}

	// Output file
	output := inputPath + ".dec"

	// Save decrypted file
	err = os.WriteFile(output, plaintext, 0644)
	if err != nil {
		return err
	}

	return nil
}