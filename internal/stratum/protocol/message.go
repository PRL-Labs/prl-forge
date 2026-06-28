package protocol

import "encoding/json"

type Request struct {
	ID     any             `json:"id"`
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
}

type Response struct {
	ID     any         `json:"id"`
	Result interface{} `json:"result"`
	Error  interface{} `json:"error"`
	Type   string      `json:"type,omitempty"`
}

type Notification struct {
	ID     any         `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}
