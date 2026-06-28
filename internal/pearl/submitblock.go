package pearl


import (

"fmt"
)

func (c *Client) SubmitBlock(block []byte) error {
	return c.Call(
		"submitblock",
		[]any{fmt.Sprintf("%x", block)},
		nil,
	)
}