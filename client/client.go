package client

import (
	"blockchain/chain"
	"time"
)

func createBlock(proof int, preHash string) chain.Block {
	block := chain.Block{
		Index:0,
		Timestamp:time.Stamp,
		Proof:proof,
	}

	return block
}
