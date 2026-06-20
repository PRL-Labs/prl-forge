package app

import (
	"log"

	"github.com/techobg/prl-forge/internal/api"
	"github.com/techobg/prl-forge/internal/api/handlers"
	"github.com/techobg/prl-forge/internal/config"
	"github.com/techobg/prl-forge/internal/pool"
	"github.com/techobg/prl-forge/internal/stratum"
)

type App struct {
	cfg      *config.Config
	api      *api.Server
	stratum  *stratum.Server
	pool     *pool.Pool
}

func New() (*App, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	poolCore := pool.New()
	handlers.SetPool(poolCore)

	return &App{
		cfg:      cfg,
		api:      api.New(cfg),
		stratum:  stratum.New(":3333"),
		pool:     poolCore,
	}, nil
}

func (a *App) Run() error {
	log.Printf("🚀 Starting %s v%s", a.cfg.App.Name, a.cfg.App.Version)

	go func() {
		if err := a.api.Start(); err != nil {
			log.Printf("API error: %v", err)
		}
	}()

	return a.stratum.Start()
}