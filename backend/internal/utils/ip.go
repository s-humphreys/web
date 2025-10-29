package utils

import (
	"net"
	"net/http"
	"strings"
)

// ClientIPFromReq handles IP extraction from a request.
func ClientIPFromReq(r *http.Request) string {
	// X-Forwarded-For: first IP is the client
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		for p := range strings.SplitSeq(xff, ",") {
			if ip := parseIP(strings.TrimSpace(p)); ip != "" {
				return ip
			}
		}
	}
	if xr := r.Header.Get("X-Real-IP"); xr != "" {
		if ip := parseIP(strings.TrimSpace(xr)); ip != "" {
			return ip
		}
	}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return parseIP(strings.TrimSpace(r.RemoteAddr))
	}
	return parseIP(host)
}

// parseIP parses and validates an IP address string.
func parseIP(s string) string {
	s = strings.Trim(s, "[]")
	if ip := net.ParseIP(s); ip != nil {
		return ip.String()
	}
	return ""
}
