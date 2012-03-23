package hello

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintln(w, "Hello, world!")

	if country, ok := r.Header["X-Appengine-Country"]; ok {
		fmt.Fprintf(w, "Your country code: %s\n", country[0])
	}
}
