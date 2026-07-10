package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

type MinerResponse struct {
	Wallet   string  `json:"wallet"`
	Hashrate float64 `json:"hashrate"`
	Workers  int     `json:"workers"`
	Status   string  `json:"status"`
}

type MinerDashboardResponse struct {
	Wallet          string    `json:"wallet"`
	CurrentHashrate float64   `json:"currentHashrate"`
	AverageHashrate float64   `json:"averageHashrate"`
	WorkersOnline   int       `json:"workersOnline"`
	LastSeen        time.Time `json:"lastSeen"`
	CurrentRound    uint64    `json:"currentRound"`
	LastBlock       string    `json:"lastBlock"`
	PersonalLuck    float64   `json:"personalLuck"`
	Blocks24h       int       `json:"blocks24h"`
	TotalBlocks     int       `json:"totalBlocks"`
}

func Miners(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if Pool == nil {
		json.NewEncoder(w).Encode([]MinerResponse{})
		return
	}

	type stat struct {
		hashrate float64
		workers  int
		online   bool
	}

	stats := make(map[string]*stat)

	for _, worker := range Pool.Workers().List() {

		s, ok := stats[worker.Wallet]
		if !ok {
			s = &stat{}
			stats[worker.Wallet] = s
		}

		s.workers++

		if time.Since(worker.LastSeen) < 2*time.Minute {
			s.hashrate += worker.Hashrate
			s.online = true
		}
	}

	resp := make([]MinerResponse, 0, len(stats))

	for wallet, s := range stats {

		status := "Offline"
		if s.online {
			status = "Online"
		}

		resp = append(resp, MinerResponse{
			Wallet:   wallet,
			Hashrate: s.hashrate,
			Workers:  s.workers,
			Status:   status,
		})
	}

	json.NewEncoder(w).Encode(resp)
}

func Miner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if Pool == nil {
		http.Error(w, "pool unavailable", http.StatusServiceUnavailable)
		return
	}

	wallet := r.URL.Query().Get("wallet")
	if wallet == "" {
		http.Error(w, "missing wallet", http.StatusBadRequest)
		return
	}

	var (
		hashrate float64
		online   int
		lastSeen time.Time
	)

	for _, worker := range Pool.Workers().List() {

		if worker.Wallet != wallet {
			continue
		}

		if time.Since(worker.LastSeen) < 2*time.Minute {
			hashrate += worker.Hashrate
			online++
		}

		if worker.LastSeen.After(lastSeen) {
			lastSeen = worker.LastSeen
		}
	}

	json.NewEncoder(w).Encode(MinerDashboardResponse{
		Wallet:          wallet,
		CurrentHashrate: hashrate,
		AverageHashrate: hashrate,
		WorkersOnline:   online,
		LastSeen:        lastSeen,
		LastBlock:       "N/A",
		PersonalLuck:    0,
		Blocks24h:       0,
		TotalBlocks:     0,
	})
}
