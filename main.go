package main

import (
	"log"

	"example.com/url_shortener/router"
)

func main() {
	// server start
	err := router.Start()
	if err != nil {
		log.Fatal(err)
	}

}
