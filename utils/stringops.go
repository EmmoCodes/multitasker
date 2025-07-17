package utils

import (
	"errors"
	"strings"
)

// takes output from validateURL, trims it and joins it together to shortURL
// returns shortURL, userInput and nil
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
	// start at 3 that it slices the first part of the URL
	// only gets the path segments
	//TODO: change it later for a rand generated id !
	for i := 3; i < len(urlSlice); i++ {
		urlSlice[i] = urlSlice[i] + "/"
	}

	urlSlice = urlSlice[3:]
	joinedURL := strings.Join(urlSlice, "")
	return joinedURL

}
