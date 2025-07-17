package main

import (
	"fmt"
	"log"

	"example.com/url_shortener/handler"
)

func main() {

	str, err := handler.New()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)

	// http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello!")
	// })

}
