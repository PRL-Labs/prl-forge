package handlers

import (
	"encoding/json"
	"net/http"
)

func Blocks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if Pool == nil {
		json.NewEncoder(w).Encode([]any{})
		return
	}

	json.NewEncoder(w).Encode(Pool.Blocks().List())
}