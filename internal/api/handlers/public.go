package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/techobg/prl-forge/internal/pool"
	"github.com/techobg/prl-forge/internal/stats"
)

type PublicStatsResponse struct {
	PoolName                string  `json:"pool_name"`
	Coin                    string  `json:"coin"`
	Algorithm               string  `json:"algorithm"`
	URL                     string  `json:"url"`
	APIURL                  string  `json:"api_url"`
	StratumHost             string  `json:"stratum_host"`
	StratumPort             int     `json:"stratum_port"`
	Fee                     float64 `json:"fee"`
	Hashrate                float64 `json:"hashrate"`
	HashrateRawHPS          float64 `json:"hashrate_raw_hps"`
	Workers                 int     `json:"workers"`
	Miners                  int     `json:"miners"`
	BlockHeight             int64   `json:"block_height"`
	NetworkHashrateEstimate float64 `json:"network_hashrate_estimate"`
	Status                  string  `json:"status"`
}

func PublicStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var (
		height            int64
		networkHashrate   float64
		poolHashrate      float64
		workers           int
	)

	tpl := stats.Template()
	if tpl != nil {
		height = tpl.Height
	}

	if Pool != nil {
		poolHashrate = float64(Pool.TotalHashrate())
		workers = Pool.OnlineWorkers()
	}

	if pool.Client() != nil {
		if hr, err := pool.Client().GetNetworkHashrate(); err == nil {
			networkHashrate = hr
		}
	}

	miners := make(map[string]struct{})

	if Pool != nil {
		for _, worker := range Pool.Workers().List() {
			miners[worker.Wallet] = struct{}{}
		}
	}

	resp := PublicStatsResponse{
		PoolName:                "PRL Forge",
		Coin:                    "PRL",
		Algorithm:               "pearlhash",
		URL:                     "https://prlforge.com",
		APIURL:                  "https://prlforge.com/api/v1/public",
		StratumHost:             "pool.prlforge.com",
		StratumPort:             3333,
		Fee:                     0.9,
		Hashrate:                poolHashrate,
		HashrateRawHPS:          poolHashrate,
		Workers:                 workers,
		Miners:                  len(miners),
		BlockHeight:             height,
		NetworkHashrateEstimate: networkHashrate,
		Status:                  "online",
	}

	json.NewEncoder(w).Encode(resp)
}