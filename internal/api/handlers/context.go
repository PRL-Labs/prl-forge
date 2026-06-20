package handlers

import "github.com/techobg/prl-forge/internal/pool"

var Pool *pool.Pool

func SetPool(p *pool.Pool) {
	Pool = p
}