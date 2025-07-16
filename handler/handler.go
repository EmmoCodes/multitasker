package handler

import (
	"net/url"
	"strings"
	"time"
)

type URLinfo struct {
	OriginalURL string
	CreatedAt   time.Time
	Clicks      int
}

type ShortURL struct {
	ShortURL  string
	CreatedAt time.Time
	URLinfo
}

func validateURL(input string) error {
	_, err := url.ParseRequestURI(input)
	if err != nil {
		panic(err)
	}
	return nil
}

func TrimURL(input string) (string, error) {
	input = "https://www.google.com/gopher/nice/jo/hallo/heyo"

	err := validateURL(input)
	if err != nil {
		panic(err)
	}

	trimmedURL := strings.Split(input, "/")
	shortURL := joinURL(trimmedURL)

	return shortURL, nil
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
