package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/techobg/prl-forge/internal/pool"
)

func MinerHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	wallet := r.URL.Query().Get("wallet")
	if wallet == "" {
		http.Error(w, "missing wallet", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(
		pool.MinerHistory().Get(wallet),
	)
}
