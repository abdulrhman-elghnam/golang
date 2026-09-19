package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
	"os"
)

func Encrypt(text string) (string, error) {
	key, err := getKey()
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	encrypted := gcm.Seal(
		nonce,
		nonce,
		[]byte(text),
		nil,
	)

	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func getKey() ([]byte, error) {
	keyHex := os.Getenv("ENCRYPTION_SECRET_KEY")
	if keyHex == "" {
		keyHex = os.Getenv("ENCRYPTION_KEY")
	}

	key, err := hex.DecodeString(keyHex)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func Decrypt(encryptedText string) (string, error) {
	key, err := getKey()
	if err != nil {
		return "", err
	}

	data, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()

	if len(data) < nonceSize {
		return "", errors.New("invalid encrypted data")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plaintext, err := gcm.Open(
		nil,
		nonce,
		ciphertext,
		nil,
	)

	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
