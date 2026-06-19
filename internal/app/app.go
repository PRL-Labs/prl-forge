package app

import (
	"github.com/techobg/prl-forge/internal/api"
)

type App struct {
	api *api.Server
}

func New() *App {
	return &App{
		api: api.New(),
	}
}

func (a *App) Run() error {
	return a.api.Start()
}