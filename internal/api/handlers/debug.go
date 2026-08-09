package handlers

import (
	"net"
	"net/http"
	"os"

	"github.com/techobg/prl-forge/internal/updater/blocks"
)

func DebugAddBlock(w http.ResponseWriter, r *http.Request) {

	// ?? allow ONLY localhost
	host, _, _ := net.SplitHostPort(r.RemoteAddr)
	if host != "127.0.0.1" && host != "::1" {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	// ?? optional env protection
	if os.Getenv("DEBUG_MODE") != "true" {
		http.Error(w, "debug disabled", http.StatusForbidden)
		return
	}

	if Pool == nil {
		http.Error(w, "pool unavailable", 500)
		return
	}

	Pool.Blocks().Add(&blocks.Block{
		Height: 999999,
		Hash:   "DEBUG_HASH",
		
	})

	w.Write([]byte("OK"))
}