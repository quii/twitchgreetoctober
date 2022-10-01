package main

import (
	"log"
	"net/http"
	"twitchgreetoctober"
)

func main() {
	if err := http.ListenAndServe(":8080", twitchgreetoctober.Handler); err != nil {
		log.Fatal(err)
	}
}
