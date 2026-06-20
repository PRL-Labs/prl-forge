package pearl

type BlockTemplate struct {
	Version           int    `json:"version"`
	PreviousBlockHash string `json:"previousblockhash"`
	Bits              string `json:"bits"`
	CurTime           int64  `json:"curtime"`
	Height            int64  `json:"height"`
	CoinbaseValue     uint64 `json:"coinbasevalue"`
}

func (c *Client) GetBlockTemplate() (*BlockTemplate, error) {

	var tpl BlockTemplate

	params := []map[string]any{
		{
			"rules": []string{},
		},
	}

	err := c.Call(
		"getblocktemplate",
		params,
		&tpl,
	)

	if err != nil {
		return nil, err
	}

	return &tpl, nil
}