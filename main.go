package main

import (
	"fmt"

	"example.com/url_shortener/handler"
)

func main() {
	val, err := handler.TrimURL("hallo")
	if err != nil {
		return
	}

	fmt.Println(val)
}
