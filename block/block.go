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
}

func NewBlock(str string, preHash []byte) *Block {
	block := &Block{
		Timestamp:time.Now().Unix(),
		PreHash:preHash,
		Hash:[]byte{},
		Data:[]byte(str),
	}

	block.buildHash()

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
