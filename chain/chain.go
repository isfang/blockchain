package chain

import "blockchain/block"

type BlockChain struct {
	blocks []*block.Block
}

// start with genesis block
func NewBlockChain() *BlockChain  {
	return &BlockChain{blocks:[]*block.Block{block.NewGenesisBlock()}}
}

func (b *BlockChain)AddBlock(str string)  {

	preBlock := b.blocks[len(b.blocks)-1]

	targetBlock := block.NewBlock(str, preBlock.Hash)

	b.blocks = append(b.blocks, targetBlock)
}

