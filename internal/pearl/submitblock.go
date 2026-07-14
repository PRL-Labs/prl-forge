package pearl

import (
	"fmt"
)

func (c *Client) SubmitBlock(block []byte) error {
	var result *string

	err := c.Call(
		"submitblock",
		[]any{fmt.Sprintf("%x", block)},
		&result,
	)
	if err != nil {
		return err
	}

	// BIP22:
	// result == nil => block accepted
	// result != nil => duplicate / rejected / high-hash / etc.
	if result != nil {
		return fmt.Errorf("submitblock rejected: %s", *result)
	}

	return nil
}
