package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"example.com/url_shortener/handler"
	"example.com/url_shortener/router"
)

func main() {
	go func() {
		// server start
		token, err := router.Start()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(token)
	}()

	time.Sleep(time.Second * 1)
	for {
		var userInput string
		choice := handler.ChoiceHandler(userInput)
		url := "http://localhost:8080" + choice
		_, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Second * 1)
	}

}
