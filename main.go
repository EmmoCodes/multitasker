package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"example.com/url_shortener/handler"
	"example.com/url_shortener/router"
)

func main() {
	//generel handle funcs for routes
	// http.HandleFunc("/", getData)
	http.HandleFunc("/create", postData)

	// server start
	err := router.Start()
	if err != nil {
		log.Fatal(err)
	}

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

// func getData(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("got /data request")
// 	io.WriteString(w, "alo")
// }
