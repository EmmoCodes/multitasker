package user

import (
	"database/sql"
	"fmt"

	"github.com/gofrs/uuid"
	_ "modernc.org/sqlite"
)

type User struct {
	Id           uuid.UUID `json:"id"`
	UserName     string    `json:"user_name"`
	UserPassword string
}

var user User

func New() error {
	var userName, userPassword, userChoice string

	fmt.Println("Please enter your username")
	fmt.Scanln(&userName)
	fmt.Printf("Thanks. you choosed : '%v' as username. Please enter now your password.\n", userName)
	fmt.Scanln(&userPassword)
	fmt.Println("Thanks. You wish to create that account now?\nPlease enter: [y/n]")
	fmt.Scanln(&userChoice)
	switch userChoice {
	case "y":
		id, err := uuid.NewV4()
		if err != nil {
			return nil
		}
		user = User{
			Id:           id,
			UserName:     userName,
			UserPassword: userPassword,
		}
		WriteUserToDb(user)
		return nil
	case "n":
		{
			fmt.Println("You choosed 'no'. If you wish create a new user now.")
		}
	}

	return nil
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

	_, err = db.Exec(`INSERT INTO users (user_name, password) VALUES (?, ?)`, user.UserName, user.UserPassword)
	if err != nil {
		return fmt.Errorf("failed to insert user into table: %w", err)
	}
	fmt.Println("User created:", user.UserName)
	return nil
}
