package pearl

func (c *Client) SubmitBlock(blockHex string) error {
	return c.Call(
		"submitblock",
		[]string{blockHex},
		nil,
	)
}