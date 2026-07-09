package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/techobg/prl-forge/internal/pool"
)

func PoolHistory(w http.ResponseWriter, r *http.Request) {

	if pool.Current() == nil {
		http.Error(w, "pool not initialized", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(
		pool.Current().History().Get(),
	)
}
