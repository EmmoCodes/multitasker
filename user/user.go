package user

import (
	"database/sql"
	"fmt"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
	_ "modernc.org/sqlite"
)

type User struct {
	Id           uuid.UUID `json:"id"`
	UserName     string    `json:"user_name"`
	UserPassword string
}

func WriteUserToDb(user User) error {
	db, err := sql.Open("sqlite", "./users.db")
	if err != nil {
		return fmt.Errorf("Failed to open file. \n%v", err)
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_name TEXT NOT NULL UNIQUE,
            password TEXT NOT NULL
        );`)
	if err != nil {
		return fmt.Errorf("Table creation failed. \n%v", err)
	}

	pwHashed, err := HashPassword(user.UserPassword)
	if err != nil {
		return err
	}
	fmt.Println("pw---->", pwHashed)
	_, err = db.Exec(`INSERT INTO users (user_name, password) VALUES (?, ?)`, user.UserName, pwHashed)
	if err != nil {
		return fmt.Errorf("failed to insert user into table: %w", err)
	}

	fmt.Println("User created:", user.UserName)
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func AuthUser(user User) error {

	db, err := sql.Open("sqlite", "./users.db")
	if err != nil {
		return fmt.Errorf("Failed to open file. \n%v", err)
	}
	db.Close()
	return nil
}
