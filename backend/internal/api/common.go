package api

import (
	"log"
	"net/http"
)

// NotFound responds with a 404 Not Found.
func NotFound(w http.ResponseWriter, r *http.Request) {
	log.Printf("not found: %s %s", r.Method, r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	_, _ = w.Write([]byte(`{"error":"not found"}`))
}
