package router

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"

	"example.com/url_shortener/handler"
	_ "modernc.org/sqlite"
)

var port string = "8080"

func Start() error {

	// general handle funcs for routes
	http.HandleFunc("/get", getData)
	http.HandleFunc("/create", postData)

	fmt.Printf("Starting server and listen to port: %v\n", port)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	return nil
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

// get data from db file via sqlite3
func getData(w http.ResponseWriter, r *http.Request) {
	// server message
	fmt.Println("got /data request")
	db, err := sql.Open("sqlite", "./urls.db")
	if err != nil {
		return
	}

	rows, err := db.Query(`SELECT short, long FROM urls;`)
	if err != nil {
		log.Fatal(err)
	}

	columns, err := rows.Columns()

	for rows.Next() {
		values := make([]string, len(columns))
		valuePtrs := make([]any, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		_ = rows.Scan(valuePtrs...)

		// printing out original and short URL
		fmt.Printf("short_url: %v\noriginal_url: %v", values[0], values[1])
		// just the short URL
		io.WriteString(w, values[0])
	}
	defer db.Close()
}
