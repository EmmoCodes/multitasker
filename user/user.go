package user

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"example.com/url_shortener/auth"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
	_ "modernc.org/sqlite"
)

type User struct {
	Id           uuid.UUID `json:"id"`
	UserName     string    `json:"user_name"`
	UserPassword string
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

func NewUser() error {
	var userName, userPassword, userChoice string

	var u User
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
			return err
		}
		u = User{
			Id:           id,
			UserName:     userName,
			UserPassword: userPassword,
		}
		WriteUserToDb(u)
		return nil
	case "n":
		{
			fmt.Println("You choosed 'no'. If you wish create a new user now.")
		}
	}

	resp, err := http.Post("http://localhost:8080/create", "application/json", strings.NewReader(`{}`))
	if err != nil {
		return err
	}
	fmt.Println(resp)
	return nil
}

func WriteUserToDb(user User) error {
	db, err := sql.Open("sqlite", "./users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_name TEXT NOT NULL UNIQUE,
            password TEXT NOT NULL
        );`)
	if err != nil {
		return err
	}

	pwHashed, err := auth.HashPassword(user.UserPassword)
	if err != nil {
		return err
	}
	_, err = db.Exec(`INSERT INTO users (user_name, password) VALUES (?, ?)`, user.UserName, pwHashed)
	if err != nil {
		return err
	}

	fmt.Println("User created:", user.UserName)
	return nil
}

// get user from db file via sqlite3
func GetUser() (Credentials, error) {
	var userName, password, storedHash string
	var creds Credentials

	fmt.Println("Please enter your username: ")
	fmt.Scanln(&userName)
	fmt.Println("Please enter your password: ")
	fmt.Scanln(&password)
	db, err := sql.Open("sqlite", "./users.db")
	if err != nil {
		return Credentials{}, err
	}

	// find username in db
	db.Exec("SELECT 1 FROM users where user = %s", userName)
	db.Exec("SELECT 1 FROM users where (user_name = %s)", userName)
	query := "SELECT 1 FROM users WHERE user_name = ?"
	row := db.QueryRow(query, userName)

	// checks if username exists
	var exists int
	err = row.Scan(&exists)
	if err != nil {
		return Credentials{}, err
	} else {
		// scan for pw and compare hashedPw(db) == inputPw
		query = "SELECT password FROM users WHERE user_name = ?"
		err = db.QueryRow(query, userName).Scan(&storedHash)
		if err != nil {
			return Credentials{}, err
		}

		err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
		if err != nil {
			fmt.Println("❌ Login failed.")
		} else {
			fmt.Println("✅ Welcome!")
			// return creds for auth func
			creds = Credentials{
				Username: userName,
				Password: password,
			}
		}
	}

	defer db.Close()
	return creds, nil
}
