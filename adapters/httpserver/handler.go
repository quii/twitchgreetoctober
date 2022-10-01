package httpserver

import (
	"fmt"
	"net/http"
	"twitchgreetoctober"
)

var Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, twitchgreetoctober.Greet(name))
})
