package configure

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/navilg/hey-jira/internal/crypt"
	"github.com/navilg/hey-jira/internal/encode"
	"github.com/navilg/hey-jira/internal/prompt"
)

type ProfileConfigurationData struct {
	Username string `json:"username"`
	B64Token string `json:"token"`
	Server   string `json:"server"`
	Project  string `json:"project"`
	// ProxyURL       string `json:"proxy_url"`
	// ProxyUsername  string `json:"proxy_username"`
	// ProxyPassword  string `json:"proxy_password"`
}

type ProfileConfiguration struct {
	Encryption        bool                     `json:"encryption"`
	EncryptionSalt    string                   `json:"encryption_salt"`
	EncryptedData     string                   `json:"encrypted_data"`
	ConfigurationData ProfileConfigurationData `json:"configuration_data"`
}

func ConfigureProfile(profile string, encryptProfile bool, encryptionPassword string) error {

	var configurationData ProfileConfigurationData
	var configuration ProfileConfiguration

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Failed to get user's home directory")
		return err
	}
	profilePath := filepath.Join(homeDir, ".hey-jira", profile)
	profileFile := filepath.Join(profilePath, "profile.json")
	if os.Stat(profileFile); err == nil {
		fmt.Printf("Profile '%s' already exist. ", profile)
		overwrite, err := prompt.GetData("Do you want to overwrite the profile (Y/n)")
		if err != nil {
			return err
		}
		if !(overwrite == "Y" || overwrite == "y") {
			fmt.Println("Aborted.")
			return nil
		}
	}
	_ = os.MkdirAll(profilePath, 0760)

	server, err := prompt.GetData("Server URL (e.g.: myjira.example.com)")
	if err != nil {
		return err
	}
	username, err := prompt.GetData("Jira username")
	if err != nil {
		return err
	}
	token, err := prompt.GetData("Jira API token")
	if err != nil {
		return err
	}
	project, err := prompt.GetData("Default project")
	if err != nil {
		return err
	}

	configurationData.Server = server
	configurationData.Username = username
	configurationData.Project = project
	configurationData.B64Token = *encode.B64Encode(&token)

	if encryptProfile == true {

		configuration.Encryption = true

		if encryptionPassword == "" {
			fmt.Println()
			encryptionPassword, err := prompt.GetPassword("Enter profile encryption password")
			if err != nil {
				return err
			}
			if encryptionPassword == "" {
				return errors.New("Password cannot be empty")
			}
		}

		configurationDataJsonByte, err := json.Marshal(configurationData)
		if err != nil {
			return err
		}

		encryptedConfigurationData, salt, err := crypt.EncryptProfile(string(configurationDataJsonByte), encryptionPassword)
		if err != nil {
			return err
		}

		configuration.EncryptedData = *encode.B64Encode(encryptedConfigurationData)
		configuration.EncryptionSalt = *encode.B64Encode(salt)

		configurationJsonByte, err := json.MarshalIndent(configuration, "", "  ")
		if err != nil {
			return err
		}

		err = os.WriteFile(profileFile, configurationJsonByte, 0760)
		if err != nil {
			fmt.Println("Failed to create profile")
			return err
		}

	} else {
		configuration.Encryption = false
		configuration.ConfigurationData = configurationData

		configurationJsonByte, err := json.MarshalIndent(configuration, "", "  ")
		if err != nil {
			return err
		}

		err = os.WriteFile(profileFile, configurationJsonByte, 0760)
		if err != nil {
			fmt.Println("Failed to create profile")
			return err
		}
	}

	fmt.Printf("profile '%s' configured\n", profile)

	return nil
}
