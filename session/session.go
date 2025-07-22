package session

import (
	"fmt"
	"time"

	"example.com/url_shortener/user"
	"github.com/gofrs/uuid"
)

var sessions = map[uuid.UUID]session{}

type session struct {
	username string
	expiry   time.Time
}

func Signin() (string, error) {
	creds, err := user.GetUser()
	if err != nil {
		return "", err
	}

	// Create a new random session token
	// we use the "github.com/google/uuid" library to generate UUIDs
	sessionToken, _ := uuid.NewV4()
	sessionTokenStr := sessionToken.String()
	expiresAt := time.Now().Add(120 * time.Second)

	// Set the token in the session map, along with the session information
	sessions[sessionToken] = session{
		username: creds.Username,
		expiry:   expiresAt,
	}

	// Finally, we set the client cookie for "session_token" as the session token we just generated
	// we also set an expiry time of 120 seconds
	// http.SetCookie(w, &http.Cookie{
	// 	Name:    "session_token",
	// 	Value:   sessionTokenStr,
	// 	Expires: expiresAt,
	// })
	fmt.Println("Session token generated.")
	return sessionTokenStr, nil
}
