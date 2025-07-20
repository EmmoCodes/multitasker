package utils

import (
	"errors"
	"fmt"
	"net/url"
)

// checks if userInput is a valid URL and returns it or nil
func ValidateURL() (string, error) {
	var userInput string

	fmt.Println("Please enter your URL to shorten:\n ")
	fmt.Scanln(&userInput)

	_, err := url.ParseRequestURI(userInput)
	if err != nil {
		return "", errors.New("Failed to validate string.\n")
	}
	return userInput, nil
}
