package crypt

import "crypto/rand"

func GenerateRandomString(size int) (string, error) {
	token := make([]byte, size)

	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	return string(token), nil
}
