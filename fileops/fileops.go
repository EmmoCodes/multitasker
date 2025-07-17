package fileops

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
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

func WriteToJson(shortenedURL ShortURL) error {
	// TODO: change it later for a db!
	fileName := "data.json"
	var urls []ShortURL

	// checks if file exists, if not it will be created
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

	// converts string to json
	urls = append(urls, shortenedURL)
	newData, err := json.MarshalIndent(urls, " ", "")
	if err != nil {
		return errors.New("Failed to marshal")
	}

	//writes json to file
	err = os.WriteFile(fileName, newData, 0644)
	if err != nil {
		return errors.New("Failed to write to file.")
	}

	fmt.Println("***Created.***")
	return nil
}
