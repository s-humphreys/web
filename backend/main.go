package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/s-humphreys/web/backend/internal/api"
	"github.com/s-humphreys/web/backend/internal/spa"
	"github.com/s-humphreys/web/backend/internal/utils"
)

//go:embed all:web/dist
var embeddedStatic embed.FS

func main() {
	root := http.NewServeMux()

	// API submux
	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/api/v1/ip", api.IPHandler)
	apiMux.HandleFunc("/api/", api.NotFound)

	// Wrap /api/* with rate limit
	root.Handle("/api/", chain(apiMux,
		utils.RateLimit(60, time.Minute), // 60 req/min per IP
	))

	// Static SPA
	dist, _ := fs.Sub(embeddedStatic, "web/dist")
	root.HandleFunc("/", spa.Handler(dist, spa.Options{
		Index:              "index.html",
		CacheControlIndex:  "no-cache",
		CacheControlAssets: "public, max-age=31536000, immutable",
	}))

	host := ":8080"
	if v := os.Getenv("PORT"); v != "" {
		host = ":" + v
	}

	log.Printf("starting server on %s", host)
	if err := http.ListenAndServe(host, root); err != nil {
		log.Fatal(err)
	}
}

// Middleware chaining
func chain(h http.Handler, m ...func(http.Handler) http.Handler) http.Handler {
	for i := len(m) - 1; i >= 0; i-- {
		h = m[i](h)
	}
	return h
}
