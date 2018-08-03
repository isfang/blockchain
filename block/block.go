package block

import (
	"time"
	"strconv"
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Timestamp int64
	PreHash []byte
	Hash []byte
	Data []byte
	Nonce int
}

func NewBlock(str string, preHash []byte) *Block {
	block := &Block{
		Timestamp:time.Now().Unix(),
		PreHash:preHash,
		Hash:[]byte{},
		Data:[]byte(str),
		Nonce:0,
	}

	//使用简单的hash创建
	//block.buildHash()

	//使用工作量证明创建hash
	pow := NewWorkProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func (self *Block) buildHash()  {
	timeStamp := []byte(strconv.FormatInt(self.Timestamp, 10))

	headers := bytes.Join([][]byte{self.PreHash, self.Data, timeStamp}, []byte{})

	hash := sha256.Sum256(headers)

	fmt.Println( "hash:", hash)

	self.Hash = hash[:]
}

func NewGenesisBlock() *Block {
	return NewBlock("genesis block", []byte{})
}
