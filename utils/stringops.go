package utils

import (
	"errors"
	"strings"
)

func TrimURL() (string, string, error) {
	userInput, err := ValidateURL()
	if err != nil {
		return "", "", errors.New("Failed to shorten URL.")
	}

	trimmedURL := strings.Split(userInput, "/")
	shortURL := joinURL(trimmedURL)
	return shortURL, userInput, nil
}

func joinURL(trimmedURL []string) string {
	var urlSlice []string

	urlSlice = append(urlSlice, trimmedURL...)
	for i := 3; i < len(urlSlice); i++ {
		urlSlice[i] = urlSlice[i] + "/"
	}

	urlSlice = urlSlice[3:]
	joinedURL := strings.Join(urlSlice, "")
	return joinedURL

}
