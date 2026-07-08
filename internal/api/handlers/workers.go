package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

type WorkerResponse struct {
	ID          string  `json:"id"`
	Wallet      string  `json:"wallet"`
	Name        string  `json:"name"`
	IP          string  `json:"ip"`
	Hashrate    float64 `json:"hashrate"`
	Shares      uint64  `json:"shares"`
	LastSeen    string  `json:"lastSeen"`
	ConnectedAt string  `json:"connectedAt"`
	Status      string  `json:"status"`
}

func Workers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if Pool == nil {
		json.NewEncoder(w).Encode([]WorkerResponse{})
		return
	}

	list := Pool.Workers().List()

	resp := make([]WorkerResponse, 0, len(list))

	for _, worker := range list {

		status := "Offline"
		if time.Since(worker.LastSeen) < 2*time.Minute {
			status = "Online"
		}

		resp = append(resp, WorkerResponse{
			ID:          worker.ID,
			Wallet:      worker.Wallet,
			Name:        worker.Name,
			IP:          worker.IP,
			Hashrate:    worker.Hashrate,
			Shares:      worker.Shares,
			LastSeen:    worker.LastSeen.Format(time.RFC3339),
			ConnectedAt: worker.ConnectedAt.Format(time.RFC3339),
			Status:      status,
		})
	}

	json.NewEncoder(w).Encode(resp)
}