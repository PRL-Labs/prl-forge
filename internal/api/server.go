package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/techobg/prl-forge/internal/api/handlers"
	"github.com/techobg/prl-forge/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func New(cfg *config.Config) *Server {
	mux := http.NewServeMux()

	// CORS Middleware
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		mux.ServeHTTP(w, r)
	})

	// Health endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, _ = w.Write([]byte(`{
	"status":"ok",
	"service":"PRL Forge",
	"version":"0.1.0"
}`))
	})

	// Dashboard endpoint
	mux.HandleFunc("/api/v1/dashboard", handlers.Dashboard)

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)

	return &Server{
		httpServer: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
	}
}

func (s *Server) Start() error {
	log.Printf("🌐 HTTP server listening on %s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("🛑 HTTP server shutting down...")
	return s.httpServer.Shutdown(ctx)
}