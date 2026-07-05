package services

import "github.com/techobg/prl-forge/internal/pearl"

type Reward struct {
	BaseReward  float64 `json:"baseReward"`
	TxFees      float64 `json:"txFees"`
	TotalReward float64 `json:"totalReward"`
	PoolFee     float64 `json:"poolFee"`
	MinerReward float64 `json:"minerReward"`
}

func CurrentReward(tpl *pearl.BlockTemplate, feePercent float64) Reward {
	var txFees float64

	for _, tx := range tpl.Transactions {
		txFees += float64(tx.Fee) / 100000000
	}

	baseReward := float64(tpl.CoinbaseValue) / 100000000

	totalReward := baseReward + txFees

	poolFee := totalReward * feePercent / 100.0

	minerReward := totalReward - poolFee

	return Reward{
		BaseReward:  baseReward,
		TxFees:      txFees,
		TotalReward: totalReward,
		PoolFee:     poolFee,
		MinerReward: minerReward,
	}
}