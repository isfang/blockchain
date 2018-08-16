package block

import (
	"fmt"
	"bytes"
	"encoding/gob"
	"crypto/sha256"
)

const subsidy = 10

type TAInput struct {
	TAId []byte 		//一个交易输入引用了之前一笔交易的一个输出, ID 表明是之前哪笔交易
	VAOut int   		//一笔交易可能有多个输出，Vout 为输出的索引
	ScriptSign string  	//提供解锁输出的签名
}

type TAOutput struct {
	Value int  					//交易数据
	ScriptPublicSign string 	//公钥
}

type Transaction struct {
	ID []byte				//交易ID 根据整个 Transcation
	VIn []TAInput			//输出
	VOut []TAOutput			//输入
}

func (ta *Transaction)SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(ta)

	if err != nil {
		fmt.Println("encode error when ta setid")
		panic(err)
	}

	hash = sha256.Sum256(encoded.Bytes())

	ta.ID = hash[:]
}


//创建 一个创世交易 也就是 新增链时的第一个
func NewGenesisTA(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("reward to '%s'", to)
	}

	tain := TAInput{[]byte{}, -1, data}
	taout := TAOutput{subsidy, to}

	ta := Transaction{nil, []TAInput{tain}, []TAOutput{taout}}

	ta.SetID()

	return &ta
}

//创建一个交易 UTXO未使用的交易输出
func NewUTXOTransaction(from, to, amount string, b *BlockChain) *Transaction {
	var ins []TAInput
	var outs []TAOutput

	//遍历链找到我的全部余额

}

