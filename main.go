package main

import (
	"example.com/url_shortener/handler"
)

func main() {
	err := handler.New()
	if err != nil {
		return
	}

}
