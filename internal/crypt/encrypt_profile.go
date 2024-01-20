package crypt

import (
	"fmt"

	"github.com/navilg/hey-jira/internal/config"
)

func EncryptProfile(profileData, encryptionPassword string) (*string, *string, error) {

	// Generate random salt string
	salt, err := GenerateRandomString(8) // 8 bytes = 64 bits of salt

	aesKey, err := GenerateKey(encryptionPassword, salt, config.AESConfig)

	if err != nil {
		fmt.Println(err.Error())
		return nil, nil, err
	}

	encryptedData, err := EncryptData(profileData, aesKey)

	// return &encryptedProfileData, &salt, error

	return encryptedData, &salt, nil
}
