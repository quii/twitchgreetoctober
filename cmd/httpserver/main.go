package main

import (
	"log"
	"net/http"
	"twitchgreetoctober/adapters/httpserver"
)

func main() {
	if err := http.ListenAndServe(":8080", httpserver.Handler); err != nil {
		log.Fatal(err)
	}
}
