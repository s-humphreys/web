package api

import (
	"encoding/json"
	"net/http"

	"github.com/s-humphreys/web/backend/internal/utils"
)

// IPHandler returns the client's IP address.
func IPHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{
		"ip": utils.ClientIPFromReq(r),
	})
}
