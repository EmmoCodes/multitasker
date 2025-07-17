package main

import (
	"log"

	"example.com/url_shortener/handler"
)

func main() {
	err := handler.New()
	if err != nil {
		log.Fatalf("Error initializing handler: %v", err)
	}

}
