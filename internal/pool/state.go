package pool

type PoolState struct {
    ActiveMiners int
    ActiveWorkers int
    AcceptedShares uint64
    RejectedShares uint64
    PoolHashrate float64
}