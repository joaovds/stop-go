package response

import (
	"encoding/json"
	"net/http"
)

func (h *HttpRes) WriteJson(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(h.StatusCode)
	json.NewEncoder(w).Encode(h)
	return
}
