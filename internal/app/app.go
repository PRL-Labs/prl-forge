package app

import (
	"log"
	"time"

	"github.com/techobg/prl-forge/internal/api"
	"github.com/techobg/prl-forge/internal/api/handlers"
	"github.com/techobg/prl-forge/internal/config"
	"github.com/techobg/prl-forge/internal/pearl"
	"github.com/techobg/prl-forge/internal/pool"
	"github.com/techobg/prl-forge/internal/stratum"
	"github.com/techobg/prl-forge/internal/updater"
    "github.com/techobg/prl-forge/internal/updater/blockwatcher"
)

type App struct {
	cfg     *config.Config
	api     *api.Server
	stratum *stratum.Server
	pool    *pool.Pool
	engine  *pool.Engine
	client  *pearl.Client
	updater *updater.Updater
}

func New() (*App, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	poolCore := pool.New()
	pool.SetCurrent(poolCore)
	handlers.SetPool(poolCore)

	engine := pool.NewEngine()
	stratum.SetEngine(engine)

	stratumServer := stratum.New(":3333")

	client := pearl.New(pearl.RPCConfig{
    Host:     cfg.Pearl.Host,
    Port:     cfg.Pearl.Port,
    User:     cfg.Pearl.User,
    Password: cfg.Pearl.Password,
})
pool.SetClient(client)
go func() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		pool.ActivityHistory().Flush()
	}
}()

	up := updater.New(
		client,
		engine,
		stratumServer,
		10*time.Second,
	)
bw := blockwatcher.New(cfg.Pool.Address)
bw.Start()
log.Println("🔥 BlockWatcher Started")
	return &App{
		cfg:     cfg,
		api:     api.New(cfg),
		stratum: stratumServer,
		pool:    poolCore,
		engine:  engine,
		client:  client,
		updater: up,
	}, nil
}

func (a *App) Run() error {
	log.Printf("🚀 Starting %s v%s", a.cfg.App.Name, a.cfg.App.Version)

	go func() {
		if err := a.api.Start(); err != nil {
			log.Printf("API error: %v", err)
		}
	}()

	a.updater.Start()

	return a.stratum.Start()
}
