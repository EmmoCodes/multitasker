package router

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"

	"example.com/url_shortener/handler"
	"example.com/url_shortener/user"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
	_ "modernc.org/sqlite"
)

var port string = "8080"

func Start() error {

	// general handle funcs for routes
	http.HandleFunc("/get", getData)
	http.HandleFunc("/new", postData)
	http.HandleFunc("/login", getUser)
	http.HandleFunc("/create", newUser)

	fmt.Printf("Starting server and listen to port: %v\n", port)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	return nil
}

// postData call a new handler to create a url slice and post it
func postData(w http.ResponseWriter, r *http.Request) {
	val, err := handler.New()
	if err != nil {
		log.Fatal(err)
	}

	// server message
	fmt.Println("got /create request")
	io.WriteString(w, val.ShortURL)
}

// get data from db file via sqlite3
func getData(w http.ResponseWriter, r *http.Request) {
	// server message
	fmt.Println("got /data request")
	db, err := sql.Open("sqlite", "./urls.db")
	if err != nil {
		return
	}

	rows, err := db.Query(`SELECT short, long FROM urls;`)
	if err != nil {
		log.Fatal(err)
	}

	columns, err := rows.Columns()

	for rows.Next() {
		values := make([]string, len(columns))
		valuePtrs := make([]any, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		_ = rows.Scan(valuePtrs...)

		// printing out original and short URL
		fmt.Printf("short_url: %v\noriginal_url: %v", values[0], values[1])
		// just the short URL
		io.WriteString(w, values[0])
	}
	defer db.Close()
}

// get user from db file via sqlite3
func getUser(w http.ResponseWriter, r *http.Request) {
	var userName, password, storedHash string

	fmt.Println("Please enter your username: ")
	fmt.Scanln(&userName)
	fmt.Println("Please enter your password: ")
	fmt.Scanln(&password)
	db, err := sql.Open("sqlite", "./users.db")
	if err != nil {
		return
	}

	pwHashed, _ := user.HashPassword(password)
	db.Exec("SELECT 1 FROM users where user = %s", userName)
	db.Exec("SELECT 1 FROM users where (user_name = %s)", userName, pwHashed)
	query := "SELECT 1 FROM users WHERE user_name = ?"
	row := db.QueryRow(query, userName)

	var exists int
	err = row.Scan(&exists)
	if err != nil {
		io.WriteString(w, "User not found. Username or password incorrect.")
		return
	} else {

		query = "SELECT password FROM users WHERE user_name = ?"
		err = db.QueryRow(query, userName).Scan(&storedHash)
		if err != nil {
			log.Println("User not found or DB error:", err)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
		if err != nil {
			fmt.Println("❌ Login failed.")
		} else {
			fmt.Println("✅ Welcome!")
		}
	}

	defer db.Close()
}

func newUser(w http.ResponseWriter, r *http.Request) {
	var userName, userPassword, userChoice string

	var u user.User
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
			return
		}
		u = user.User{
			Id:           id,
			UserName:     userName,
			UserPassword: userPassword,
		}
		user.WriteUserToDb(u)
		return
	case "n":
		{
			fmt.Println("You choosed 'no'. If you wish create a new user now.")
		}
	}
	io.WriteString(w, "User succesfull created.")
}
