package main

import (
	"fmt"
	"blockchain/block"
	"strconv"
)

func main() {

	blockChain := block.NewBlockChain()

	blockChain.AddBlock("block-1")
	blockChain.AddBlock("block-2")
	blockChain.AddBlock("block-3")

	for _, b := range  blockChain.Blocks {
		fmt.Printf("Prev hash: %x\n", b.PreHash)
		fmt.Printf("Data: %s\n", b.Data)
		fmt.Printf("Hash: %x\n", b.Hash)
		pow := block.NewWorkProof(b)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

}
