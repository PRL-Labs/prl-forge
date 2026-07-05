package pool

import "github.com/techobg/prl-forge/internal/pearl"

var rpcClient *pearl.Client
var currentPool *Pool

func SetClient(c *pearl.Client) {
	rpcClient = c
}

func Client() *pearl.Client {
	return rpcClient
}

func SetCurrent(p *Pool) {
	currentPool = p
}

func Current() *Pool {
	return currentPool
}