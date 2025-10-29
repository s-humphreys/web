package utils

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type slot struct {
	count int
	reset time.Time
}

// In-memory per-IP rate limit (not shared across instances).
func RateLimit(max int, window time.Duration) func(http.Handler) http.Handler {
	var (
		mu   sync.Mutex
		bkt  = map[string]*slot{}
		nowf = time.Now
	)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := ClientIPFromReq(r)
			now := nowf()

			mu.Lock()
			s := bkt[ip]
			if s == nil || now.After(s.reset) {
				s = &slot{count: 0, reset: now.Add(window)}
				bkt[ip] = s
			}
			s.count++
			count := s.count
			resetIn := int(time.Until(s.reset).Seconds())
			mu.Unlock()

			w.Header().Set("X-RateLimit-Limit", strconv.Itoa(max))
			w.Header().Set("X-RateLimit-Remaining", strconv.Itoa(max-count))
			w.Header().Set("X-RateLimit-Reset", strconv.Itoa(resetIn))

			if count > max {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusTooManyRequests)
				_ = json.NewEncoder(w).Encode(map[string]string{"error": "rate_limited"})
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
