package BLC

import (
	"bytes"
	"encoding/gob"
	"log"
	"crypto/sha256"
)

//定义交易的数据
type Transaction struct {
	//1.交易ID(该笔交易的Hash)
	TxID []byte
	//2.输入
	Vins []*TxInput
	//3.输出
	Vouts []*TxOutput
}

/*
交易
1.CoinBase交易:创世区块中
2.转账产生的普通交易：

 */

func NewCoinBaseTransaction(address string) *Transaction {
	//1.创建第一个交易的Input记录
	txInput := &TxInput{TxID: []byte{}, Vout: -1, ScriptSiq: "Genesis Block"}
	//2.创建第一个交易的Output记录
	txOutput := &TxOutput{Value: 10, ScriptPubKey: address}
	txCoinBaseTransaction := &Transaction{TxID: []byte{}, Vins: []*TxInput{txInput}, Vouts: []*TxOutput{txOutput}}

	//设置交易ID
	txCoinBaseTransaction.SetID()
	return txCoinBaseTransaction

}

//TxID -->根据交易生成Hash
func (tx *Transaction) SetID() {

	//1.tx ->[]byte
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	//2.[]byte ->hash
	hash := sha256.Sum256(buf.Bytes())
	//3.为tx交易对象设置ID
	tx.TxID = hash[:]

}
