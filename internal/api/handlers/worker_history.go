package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/techobg/prl-forge/internal/pool"
)

func WorkerHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	wallet := r.URL.Query().Get("wallet")
	worker := r.URL.Query().Get("worker")

	if wallet == "" || worker == "" {
		http.Error(w, "missing wallet or worker", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(
		pool.WorkerHistory().Get(wallet, worker),
	)
}
