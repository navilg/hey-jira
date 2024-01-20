package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/navilg/hey-jira/internal/config"
	"golang.org/x/crypto/scrypt"
)

func GenerateKey(password, salt string, aesconfig config.KDFConfiguration) (string, error) {
	passwordBytes := []byte(password)
	saltBytes := []byte(salt)

	derivedKey, err := scrypt.Key(passwordBytes, saltBytes, aesconfig.CostFactor, aesconfig.BlockSizeFactor, aesconfig.ParallelizationFactor, aesconfig.KeySize)

	if err != nil {
		return "", err
	}

	return string(derivedKey), nil
}

func EncryptData(message, key string) (*string, error) {

	messageBytes := []byte(message)
	keyBytes := []byte(key)

	cipherBlock, err := aes.NewCipher(keyBytes)

	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(cipherBlock)

	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize()) // Standard 12 bytes

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	encryptedMessageBytes := gcm.Seal(nonce, nonce, messageBytes, nil)

	encryptedMessageStr := string(encryptedMessageBytes)

	return &encryptedMessageStr, nil
}
