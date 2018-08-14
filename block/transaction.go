package block

import "fmt"

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
	ID []byte				//交易ID
	VIn []TAInput			//输出
	VOut []TAOutput			//输入
}


//创建 一个创世交易 也就是 新增链时的第一个
func NewGenesisTA(to, data string)  {
	if data == "" {
		data = fmt.Sprintf("reward to '%s'", to)
	}

	tain := TAInput{[]byte{}, -1, data}
	taout := TAOutput{subsidy, to}

	ta := Transaction{nil, []TAInput{tain}, []TAOutput{taout}}

	ta.set
}

™