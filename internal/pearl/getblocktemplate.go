package pearl

type Transaction struct {
	Data string `json:"data"`
	Hash string `json:"hash"`
	TxID string `json:"txid"`
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

	CoinbaseValue uint64        `json:"coinbasevalue"`
	CoinbaseAux   CoinbaseAux   `json:"coinbaseaux"`
	Transactions  []Transaction `json:"transactions"`
RequiredCertVersion uint32 `json:"requiredcertversion"`
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
println("RequiredCertVersion:", tpl.RequiredCertVersion)
	return &tpl, nil
}
func (c *Client) GetDifficulty() (float64, error) {
	var difficulty float64

	if err := c.Call(
		"getdifficulty",
		[]any{},
		&difficulty,
	); err != nil {
		return 0, err
	}

	return difficulty, nil
}

func (c *Client) GetNetworkHashrate() (float64, error) {
	var hashrate float64

	if err := c.Call(
		"getnetworkhashps",
		[]any{},
		&hashrate,
	); err != nil {
		return 0, err
	}

	return hashrate, nil
}