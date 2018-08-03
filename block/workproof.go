package block

import (
	"math/big"
	"math"
	"fmt"
	"bytes"
	"blockchain/util"
	"crypto/sha256"
)

const targetBits = 24
const maxNonce = math.MaxInt64

type WorkProof struct {
	block *Block
	target *big.Int
}

func NewWorkProof(b *Block) *WorkProof {

	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	fmt.Println("block ", string(b.Data),"target is ", target)

	pow := &WorkProof{
		block:b,
		target:target,
	}
	return pow
}

//工作量证明数据
func (p *WorkProof)prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			p.block.PreHash,
			p.block.Data,
			util.IntToHex(p.block.Timestamp),
			util.IntToHex(int64(targetBits)),
			util.IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

func (p *WorkProof) Run()(int, []byte)  {
	var hashInt big.Int
	var hash [32]byte

	nonce := 0
	fmt.Printf("Mining the block containing \"%s\"\n", p.block.Data)

	for nonce < maxNonce {
		data := p.prepareData(nonce)

		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(p.target) == -1 {
			break
		} else {
			nonce++
		}
	}

	fmt.Print("\n-------------------------------\n")

	return nonce, hash[:]
}

func (p *WorkProof)Validate()  bool {
	var hashInt big.Int

	data := p.prepareData(p.block.Nonce)
	hash := sha256.Sum256(data)

	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(p.target) == -1
}