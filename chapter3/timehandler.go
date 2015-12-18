package main

import (
	"fmt"
	"net/http"
	"time"
)

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
	http.Handle("/", NewUptimeHandler())
	http.ListenAndServe(":3000", nil)
}