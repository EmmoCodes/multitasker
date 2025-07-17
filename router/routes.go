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

func Get() error {
	Start()
	url := "https://localhost:" + port + "/hello"
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Failed to get request: %v", err)
	}
	fmt.Println(resp)
	return nil
}
