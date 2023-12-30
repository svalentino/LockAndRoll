package cmd

import (
	"fmt"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func PromptSecretValue() string {
	fmt.Printf("Enter the value of the secret: \n")
	byteSecretValue, _ := term.ReadPassword(int(syscall.Stdin))
	secretValue := strings.TrimSpace(string(byteSecretValue))

	return secretValue
}
