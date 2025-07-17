package router

import (
	"fmt"
	"log"
	"net/http"
)

var port string = "8080"

type PageData struct {
	Title    string
	ShortURL string
	URL      string
}

func Start() error {
	fmt.Printf("Starting server and listen to port: %v\n\n\n", port)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	return nil
}
