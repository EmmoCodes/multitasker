package handler

import (
	"errors"
	"time"

	"example.com/url_shortener/fileops"
	"example.com/url_shortener/router"
	"example.com/url_shortener/utils"
	"github.com/gofrs/uuid"
)

func New() (fileops.ShortURL, error) {
	go func() {
		router.Start()
	}()

	time.Sleep(time.Second * 1)
	trimmedURL, userInput, err := utils.TrimURL()

	if err != nil {
		return fileops.ShortURL{}, errors.New("Failed to get shortened URL")
	}

	id, err := uuid.NewV4()
	if err != nil {
		return fileops.ShortURL{}, errors.New("Failed to create id.")
	}

	originalURL := fileops.URLInfo{
		Id:          id,
		OriginalURL: userInput,
		CreatedAt:   time.Now(),
	}

	shortenedURL := fileops.ShortURL{
		Id:        id,
		ShortURL:  trimmedURL,
		CreatedAt: time.Now(),
		URLInfo:   originalURL,
	}

	err = fileops.WriteToJson(shortenedURL)
	if err != nil {
		return fileops.ShortURL{}, errors.New("Failed to write to json")
	}

	return shortenedURL, nil
}
