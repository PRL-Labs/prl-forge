package pearl

type Transaction struct {
	Data string `json:"data"`
	Hash string `json:"hash"`
	Fee  int64  `json:"fee"`
}

type CoinbaseAux struct {
	Flags string `json:"flags"`
}

type BlockTemplate struct {
	Version           int    `json:"version"`
	PreviousBlockHash string `json:"previousblockhash"`
	Bits              string `json:"bits"`
	CurTime           int64  `json:"curtime"`
	Height            int64  `json:"height"`

	CoinbaseValue uint64       `json:"coinbasevalue"`
	CoinbaseAux   CoinbaseAux  `json:"coinbaseaux"`
	Transactions  []Transaction `json:"transactions"`

	Target string `json:"target,omitempty"`
}

func (c *Client) GetBlockTemplate() (*BlockTemplate, error) {
	var tpl BlockTemplate

	params := []map[string]any{
		{
			"rules": []string{},
		},
	}

	if err := c.Call(
		"getblocktemplate",
		params,
		&tpl,
	); err != nil {
		return nil, err
	}

	return &tpl, nil
}