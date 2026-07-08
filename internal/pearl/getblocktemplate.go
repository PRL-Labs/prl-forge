


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

func (c *Client) GetBlockCount() (int64, error) {
	var height int64

	if err := c.Call(
		"getblockcount",
		[]any{},
		&height,
	); err != nil {
		return 0, err
	}

	return height, nil
}

func (c *Client) GetBlockHash(height int64) (string, error) {
	var hash string

	if err := c.Call(
		"getblockhash",
		[]any{height},
		&hash,
	); err != nil {
		return "", err
	}

	return hash, nil
}

type ScriptPubKey struct {
	Address string `json:"address"`
}

type Vout struct {
	Value        float64      `json:"value"`
	ScriptPubKey ScriptPubKey `json:"scriptPubKey"`
}

type RawTx struct {
	Vout []Vout `json:"vout"`
}

type Block struct {
	Hash          string  `json:"hash"`
	Height        int64   `json:"height"`
	Time          int64   `json:"time"`
	Confirmations int64   `json:"confirmations"`

	RawTx []RawTx `json:"rawtx"`
}

func (c *Client) GetBlock(hash string) (*Block, error) {

	var block Block

	if err := c.Call(
		"getblock",
		[]any{hash, 2},
		&block,
	); err != nil {
		return nil, err
	}

	return &block, nil
}