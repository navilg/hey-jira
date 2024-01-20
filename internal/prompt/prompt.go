package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func GetPassword(prompt string) (string, error) {
	// reader := bufio.NewReader(os.Stdin)

	fmt.Print(prompt + " > ")

	password, err := term.ReadPassword(int(syscall.Stdin))

	// password, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	fmt.Println()

	return string(password), nil
}

func GetData(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(prompt + " > ")
	message, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error: Failed to read your message.")
		fmt.Println(err.Error())
		return "", err
	}

	message = strings.TrimSuffix(message, "\n")

	return message, nil
}
