package api

import (
	"context"
	"log"
	"net/http"
	"github.com/techobg/prl-forge/internal/config"
	"fmt"
)

type Server struct {
	httpServer *http.Server
}

func New(cfg *config.Config) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, _ = w.Write([]byte(`{
	"status":"ok",
	"service":"PRL Forge",
	"version":"0.1.0"
}`))
	})

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)

return &Server{
	httpServer: &http.Server{
		Addr:    addr,
		Handler: mux,
	},
}
}

func (s *Server) Start() error {
	log.Println("🌐 HTTP server listening on :8080")
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("🛑 HTTP server shutting down...")
	return s.httpServer.Shutdown(ctx)
}