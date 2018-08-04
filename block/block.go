package block

import (
	"time"
	"strconv"
	"bytes"
	"crypto/sha256"
	"fmt"
	"encoding/gob"
	"log"
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

func NewGenesisBlock() *Block {
	return NewBlock("genesis block", []byte{})
}

func (b *Block) buildHash()  {
	timeStamp := []byte(strconv.FormatInt(b.Timestamp, 10))

	headers := bytes.Join([][]byte{b.PreHash, b.Data, timeStamp}, []byte{})

	hash := sha256.Sum256(headers)

	fmt.Println( "hash:", hash)

	b.Hash = hash[:]
}

func (b *Block)Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)

	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func Deserialize(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
