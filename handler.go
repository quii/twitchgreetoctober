package twitchgreetoctober

import (
	"fmt"
	"net/http"
)

var Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, Greet(name))
})
