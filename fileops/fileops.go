package fileops

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gofrs/uuid"
	_ "modernc.org/sqlite"
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

// if wished so a func to write a json file too, for own documentation of shortened URL's
func WriteToJson(shortenedURL ShortURL) error {
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

func WriteToDb(short, long string) error {
	db, err := sql.Open("sqlite", "./urls.db")
	if err != nil {
		return fmt.Errorf("Failed to open file. \n%v", err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS urls (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            short TEXT NOT NULL UNIQUE,
            long TEXT NOT NULL
        );`)
	if err != nil {
		return fmt.Errorf("Table creation failed. \n%v", err)
	}

	_, err = db.Exec(`INSERT INTO urls (short, long) VALUES (?, ?)`, short, long)
	if err != nil {
		return fmt.Errorf("failed to insert into table: %w", err)
	}
	fmt.Println("Gespeichert:", short, "→", long)
	return nil
}
