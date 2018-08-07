package block

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

type BlockChain struct {
	Tip []byte
	BlotDB *bolt.DB
}



// start with genesis block
func NewBlockChain() *BlockChain  {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)

	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			fmt.Println("create genesis block")
			genesis := NewGenesisBlock()

			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}

			b.Put(genesis.Hash, genesis.Serialize())

			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				log.Panic(err)
			}
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	bc := BlockChain{
		Tip:tip,
		BlotDB:db,
	}

	return &bc
}

func (bc *BlockChain)AddBlock(str string)  {

	//preBlock := b.Blocks[len(b.Blocks)-1]
	//
	//
	//targetBlock := NewBlock(str, preBlock.Hash)
	//
	//b.Blocks = append(b.Blocks, targetBlock)

	var lastHash []byte
	err := bc.BlotDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	nb := NewBlock(str, lastHash)

	bc.BlotDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(nb.Hash, nb.Serialize())
		if err != nil {
			log.Panic(err)
		}
		b.Put([]byte("l"), nb.Hash)
		if err != nil {
			log.Panic(err)
		}
		bc.Tip = nb.Hash

		return nil
	})
}

type BlockchainIterator struct {
	currentHash []byte
	db *bolt.DB
}

func (bc *BlockChain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.Tip, bc.BlotDB}

	return bci
}

// 返回链中的下一个块
func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = Deserialize(encodedBlock)

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	i.currentHash = block.PreHash

	return block
}


//type BlockChain struct {
//	Blocks []*Block
//}
//
//
//
//// start with genesis block
//func NewBlockChain() *BlockChain  {
//	return &BlockChain{Blocks:[]*Block{NewGenesisBlock()}}
//}
//
//func (b *BlockChain)AddBlock(str string)  {
//
//	preBlock := b.Blocks[len(b.Blocks)-1]
//
//	targetBlock := NewBlock(str, preBlock.Hash)
//
//	b.Blocks = append(b.Blocks, targetBlock)
//}

