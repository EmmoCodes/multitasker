package router

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"example.com/url_shortener/handler"
	"example.com/url_shortener/user"
	"github.com/gofrs/uuid"
	_ "modernc.org/sqlite"
)

var port string = "8080"

var sessions = map[uuid.UUID]session{}

type session struct {
	username string
	expiry   time.Time
}

func Start() (user.Credentials, error) {
	// general handle funcs for routes
	http.HandleFunc("/get", getData)
	http.HandleFunc("/new", postData)
	http.HandleFunc("/login", signin)
	http.HandleFunc("/register", newUser)

	fmt.Printf("Starting server and listen to port: %v\n", port)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	creds, err := user.GetUser()
	if err != nil {
		fmt.Println("Please login or create an account.")
	}

	return creds, nil
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

func newUser(w http.ResponseWriter, r *http.Request) {

	err := user.NewUser()
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, "Login successful.")
}

func signin(w http.ResponseWriter, r *http.Request) {

	creds, err := user.GetUser()
	if err != nil {
		log.Fatal(err)
	}
	// Get the JSON body and decode into credentials
	err = json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		// w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create a new random session token
	// we use the "github.com/google/uuid" library to generate UUIDs
	sessionToken, _ := uuid.NewV4()
	sessionTokenStr := sessionToken.String()
	expiresAt := time.Now().Add(12 * time.Second)

	// Set the token in the session map, along with the session information
	sessions[sessionToken] = session{
		username: creds.Username,
		expiry:   expiresAt,
	}

	// Finally, we set the client cookie for "session_token" as the session token we just generated
	// we also set an expiry time of 120 seconds
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionTokenStr,
		Expires: expiresAt,
	})
	fmt.Println("Session token generated.")
}
