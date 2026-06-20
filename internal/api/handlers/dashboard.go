package handlers

import (
	"encoding/json"
	"net/http"
)

type DashboardResponse struct {
	Pool    PoolInfo    `json:"pool"`
	Workers WorkersInfo `json:"workers"`
	Network NetworkInfo `json:"network"`
}

type PoolInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Status  string `json:"status"`
}

type WorkersInfo struct {
	Online int `json:"online"`
}

type NetworkInfo struct {
	Height     int64 `json:"height"`
	Difficulty int64 `json:"difficulty"`
	Hashrate   int64 `json:"hashrate"`
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	resp := DashboardResponse{
		Pool: PoolInfo{
			Name:    "PRL Forge",
			Version: "0.1.0",
			Status:  "online",
		},
		Workers: WorkersInfo{
			Online: 0,
		},
		Network: NetworkInfo{
			Height:     0,
			Difficulty: 0,
			Hashrate:   0,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}