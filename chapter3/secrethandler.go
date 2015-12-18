package main

import (
	"fmt"
	"net/http"
	"time"
)

// SecretTokenHandler secures a request with a secret token
type SecretTokenHandler struct {
	next 	http.Handler
	secret 	string
}

// ServeHTTP makes SecretTokenHandler implement http.Handler interface
func (h SecretTokenHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Check the query string for secret token
	if req.URL.Query().Get("secret_token") == h.secret {
		// secret mathced, call next handler
		h.next.ServeHTTP(w, req)
	}else{
		// no match, return 404
		http.NotFound(w, req)
	}
}

// Uptime handler writes the numner of seconds since startin the response
type UptimeHandler struct {
	Started time.Time
}

func NewUptimeHandler() UptimeHandler {
	return UptimeHandler{ time.Now() }
}

func (h UptimeHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(
		w,
		fmt.Sprintf("Current Uptime: %s", time.Since(h.Started)),
	)
}

func main() {
	http.Handle("/", SecretTokenHandler{
		next:		NewUptimeHandler(),
		secret:		"MySecret",
	})

	http.ListenAndServe(":3000", nil)
}