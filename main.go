package main

import (
	"blockchain/chain"
)

func main() {

	blockChain := chain.NewBlockChain()

	blockChain.AddBlock("eat hhh")
	blockChain.AddBlock("eat hhh")
	blockChain.AddBlock("eat hhh")
	blockChain.AddBlock("eat hhh")
}
