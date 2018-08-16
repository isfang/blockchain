package block

import (
	"time"
	"bytes"
	"encoding/gob"
	"log"
	"crypto/sha256"
)

type Block struct {
	Timestamp int64
	PreHash []byte
	Hash []byte
	Transactions []*Transaction
	Nonce int
}

func NewBlock(transcations []*Transaction, preHash []byte) *Block {
	block := &Block{
		Timestamp:time.Now().Unix(),
		PreHash:preHash,
		Hash:[]byte{},
		Transactions:transcations,
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

func NewGenesisBlock(ta *Transaction) *Block {
	return NewBlock([]*Transaction{ta}, []byte{})
}

//func (b *Block) buildHash()  {
//	timeStamp := []byte(strconv.FormatInt(b.Timestamp, 10))
//
//	headers := bytes.Join([][]byte{b.PreHash, b.Data, timeStamp}, []byte{})
//
//	hash := sha256.Sum256(headers)
//
//	fmt.Println( "hash:", hash)
//
//	b.Hash = hash[:]
//}

func (b *Block) buildHash() []byte {

	var tempHash [][]byte
	var resultHash [32]byte

	for _, transaction := range b.Transactions{
		tempHash = append(tempHash, transaction.ID)
	}

	resultHash = sha256.Sum256(bytes.Join(tempHash, []byte{}))

	return resultHash[:]
}

// 计算区块里所有交易的哈希
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
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
