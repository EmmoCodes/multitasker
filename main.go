package main

import (
	"example.com/url_shortener/router"
	"log"
)

func main() {
	// server start
	err := router.Start()
	if err != nil {
		log.Fatal(err)
	}

}
