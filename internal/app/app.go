package app

import (
	"log"

	"github.com/techobg/prl-forge/internal/api"
	"github.com/techobg/prl-forge/internal/config"
)

type App struct {
	cfg *config.Config
	api *api.Server
}

func New() (*App, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	return &App{
		cfg: cfg,
		api: api.New(cfg),
	}, nil
}

func (a *App) Run() error {
	log.Printf(
		"🚀 Starting %s v%s",
		a.cfg.App.Name,
		a.cfg.App.Version,
	)

	return a.api.Start()
}