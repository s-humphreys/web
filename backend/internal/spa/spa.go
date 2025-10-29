package spa

import (
	"bytes"
	"io/fs"
	"mime"
	"net/http"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type Options struct {
	Index              string // default: index.html
	CacheControlIndex  string // e.g. "no-cache"
	CacheControlAssets string // e.g. "public, max-age=31536000, immutable"
}

// Handler serves a Single Page Application (SPA) with fallback to index.html.
func Handler(static fs.FS, opts Options) http.HandlerFunc {
	index := opts.Index
	if index == "" {
		index = "index.html"
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// Never serve SPA for API paths
		if strings.HasPrefix(r.URL.Path, "/api/") {
			http.NotFound(w, r)
			return
		}

		p := strings.TrimPrefix(path.Clean(r.URL.Path), "/")
		if p == "" {
			p = index
		}

		// Try asset, else fallback to index.html
		if err := serveFSFile(w, r, static, p, opts.CacheControlAssets); err != nil {
			if err := serveFSFile(w, r, static, index, opts.CacheControlIndex); err != nil {
				http.NotFound(w, r)
			}
			return
		}
	}
}

// serveFSFile serves a file from the given fs.FS with appropriate headers.
func serveFSFile(w http.ResponseWriter, r *http.Request, static fs.FS, name, cacheControl string) error {
	data, err := fs.ReadFile(static, name)
	if err != nil {
		return err
	}

	if cacheControl != "" {
		w.Header().Set("Cache-Control", cacheControl)
	}
	if ctype := mime.TypeByExtension(filepath.Ext(name)); ctype != "" {
		w.Header().Set("Content-Type", ctype)
	}

	modTime := time.Time{}
	if info, err := fs.Stat(static, name); err == nil {
		modTime = info.ModTime()
	}

	http.ServeContent(w, r, name, modTime, bytes.NewReader(data))
	return nil
}
