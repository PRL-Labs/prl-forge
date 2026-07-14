package pool

import (
	"github.com/techobg/prl-forge/internal/history"
	"github.com/techobg/prl-forge/internal/pearl"
  "github.com/techobg/prl-forge/internal/vardiff"
)

var rpcClient *pearl.Client
var currentPool *Pool
var minerHistory = history.NewMinerManager()
var activityHistory = history.NewActivityManager()

var workerHistory = history.NewWorkerManager()
var varDiff = vardiff.New()

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

func WorkerHistory() *history.WorkerManager {
	return workerHistory
}
func MinerHistory() *history.MinerManager {
	return minerHistory
}
func ActivityHistory() *history.ActivityManager {
	return activityHistory
}

func VarDiff() *vardiff.Manager {
        return varDiff
}