package pearl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	config RPCConfig
	http   *http.Client
}

func New(config RPCConfig) *Client {
	return &Client{
		config: config,
		http:   NewHTTPClient(),
	}
}

type rpcRequest struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Method  string `json:"method"`
	Params  any    `json:"params"`
}

type rpcResponse struct {
	Result json.RawMessage `json:"result"`
	Error  any             `json:"error"`
	ID     int             `json:"id"`
}

func (c *Client) Call(method string, params any, result any) error {

	reqBody := rpcRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  method,
		Params:  params,
	}

	data, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		c.config.URL(),
		bytes.NewBuffer(data),
	)

	if err != nil {
		return err
	}

	req.SetBasicAuth(
		c.config.User,
		c.config.Password,
	)

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}



	var rpcResp rpcResponse

	if err := json.Unmarshal(body, &rpcResp); err != nil {
		return err
	}

	if rpcResp.Error != nil {
		return fmt.Errorf("rpc error: %v", rpcResp.Error)
	}

	if result != nil {
		if err := json.Unmarshal(rpcResp.Result, result); err != nil {
			return err
		}
	}

	return nil
}
