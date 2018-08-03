package block


type BlockChain struct {
	Blocks []*Block
}

// start with genesis block
func NewBlockChain() *BlockChain  {
	return &BlockChain{Blocks:[]*Block{NewGenesisBlock()}}
}

func (b *BlockChain)AddBlock(str string)  {

	preBlock := b.Blocks[len(b.Blocks)-1]

	targetBlock := NewBlock(str, preBlock.Hash)

	b.Blocks = append(b.Blocks, targetBlock)
}

