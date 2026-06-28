package pool

import "github.com/techobg/prl-forge/internal/pearl"

var rpcClient *pearl.Client

func SetClient(c *pearl.Client) {
	rpcClient = c
}