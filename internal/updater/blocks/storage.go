package blocks

import (
	"encoding/json"
	"os"
)

const blocksFile = "data/blocks.json"

func SaveBlocks(list []*Block) error {
	_ = os.MkdirAll("data", 0755)

	f, err := os.Create(blocksFile)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(list)
}

func LoadBlocks() ([]*Block, error) {
	f, err := os.Open(blocksFile)
	if err != nil {
		return []*Block{}, err
	}
	defer f.Close()

	var list []*Block

	if err := json.NewDecoder(f).Decode(&list); err != nil {
		return []*Block{}, err
	}

	return list, nil
}