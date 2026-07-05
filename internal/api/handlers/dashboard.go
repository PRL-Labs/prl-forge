package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/techobg/prl-forge/internal/pool"
	"github.com/techobg/prl-forge/internal/services"
	"github.com/techobg/prl-forge/internal/stats"
)

type DashboardResponse struct {
	Pool    PoolInfo        `json:"pool"`
	Workers WorkersInfo     `json:"workers"`
	Network NetworkInfo     `json:"network"`
	Reward  services.Reward `json:"reward"`
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
	Height          int64   `json:"height"`
	Difficulty      int64   `json:"difficulty"`
	NetworkHashrate float64 `json:"networkHashrate"`
	PoolHashrate    int64   `json:"poolHashrate"`
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	tpl := stats.Template()

	height := int64(0)
	if tpl != nil {
		height = tpl.Height
	}

	difficulty := int64(0)
	if pool.Client() != nil {
		if diff, err := pool.Client().GetDifficulty(); err == nil {
			difficulty = int64(diff)
		}
	}

	networkHashrate := float64(0)
	if pool.Client() != nil {
		if hr, err := pool.Client().GetNetworkHashrate(); err == nil {
			networkHashrate = hr
		}
	}

	reward := services.Reward{}
	if tpl != nil {
		reward = services.CurrentReward(tpl, 1.5)
	}

	resp := DashboardResponse{
		Pool: PoolInfo{
			Name:    "PRL Forge",
			Version: "0.1.0",
			Status:  "online",
		},

		Workers: WorkersInfo{
			Online: func() int {
				if Pool == nil {
					return 0
				}
				return Pool.OnlineWorkers()
			}(),
		},

		Network: NetworkInfo{
			Height:          height,
			Difficulty:      difficulty,
			NetworkHashrate: networkHashrate,
			PoolHashrate: func() int64 {
				if Pool == nil {
					return 0
				}
				return Pool.TotalHashrate()
			}(),
		},

		Reward: reward,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}