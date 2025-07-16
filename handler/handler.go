package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gofrs/uuid"
)

type URLInfo struct {
	Id          uuid.UUID `json:"id"`
	OriginalURL string    `json:"url"`
	CreatedAt   time.Time `json:"created_at"`
}

type ShortURL struct {
	Id        uuid.UUID `json:"id"`
	ShortURL  string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	URLInfo   URLInfo   `json:"original_url"`
}

func New() error {
	var wg sync.WaitGroup
	fileName := "data.json"
	wg.Add(1)

	go func() error {
		defer wg.Done()
		trimmedURL, userInput, err := TrimURL(&wg)
		if err != nil {
			panic(err)
		}

		id, err := uuid.NewV4()
		if err != nil {
			return errors.New("Failed to create id.")
		}

		originalURL := URLInfo{
			Id:          id,
			OriginalURL: userInput,
			CreatedAt:   time.Now(),
		}

		shortenedURL := ShortURL{
			Id:        id,
			ShortURL:  trimmedURL,
			CreatedAt: time.Now(),
			URLInfo:   originalURL,
		}

		var urls []ShortURL

		if _, err := os.Stat(fileName); err == nil {
			content, err := os.ReadFile(fileName)
			if err != nil {
				return errors.New("Failed to find file.")
			}

			if len(content) > 0 {
				err = json.Unmarshal(content, &urls)
				if err != nil {
					return errors.New("Failed to unmarshal .")
				}
			}
		}

		urls = append(urls, shortenedURL)
		newData, err := json.Marshal(urls)
		if err != nil {
			return errors.New("Failed to marshal")
		}

		err = os.WriteFile(fileName, newData, 0644)
		if err != nil {
			return errors.New("Failed to write to file.")
		}

		return nil
	}()
	wg.Wait()
	fmt.Println("***Created.***")
	return nil
}

func validateURL() (string, error) {
	var userInput string
	fmt.Println("Please enter your URL to shorten: ")
	fmt.Scanln(&userInput)

	_, err := url.ParseRequestURI(userInput)
	if err != nil {
		return "", errors.New("Failed to validate string.\n")
	}
	return userInput, nil
}

func TrimURL(wg *sync.WaitGroup) (string, string, error) {

	userInput, err := validateURL()
	if err != nil {
		panic(err)
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
