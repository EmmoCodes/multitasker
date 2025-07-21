package handler

import (
	"errors"
	"fmt"
	"time"

	"example.com/url_shortener/fileops"
	"example.com/url_shortener/utils"
	"github.com/gofrs/uuid"
)

func New() (fileops.ShortURL, error) {

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

	// => comment in this following lines to write a json file <=
	// err = fileops.WriteToJson(shortenedURL)
	// if err != nil {
	// 	return fileops.ShortURL{}, errors.New("Failed to write to json")
	// }

	err = fileops.WriteToDb(trimmedURL, userInput)
	if err != nil {
		return fileops.ShortURL{}, fmt.Errorf("failed writing to db: %v ", err)

	}

	return shortenedURL, nil
}

func ChoiceHandler(val string) string {
	var choice string

	fmt.Println("Please select one of the following options: ")
	fmt.Println("1: Register")
	fmt.Println("2: Login")
	fmt.Println("3: New URL")
	fmt.Println("4: Show URL list")
	fmt.Println("q: Exit application")
	fmt.Scan(&val)

	switch val {
	case "1":
		choice = "/register"
	case "2":
		choice = "/login"
	case "3":
		choice = "/new"
	case "4":
		choice = "/get"
	case "q":
		choice = "/exit"
	default:
		fmt.Println("Please choose an option.")
	}
	return choice
}
