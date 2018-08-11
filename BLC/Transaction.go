package BLC

import (
	"bytes"
	"encoding/gob"
	"log"
	"crypto/sha256"
	"encoding/hex"
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

//根据转账信息，创建一个普通的信息
func NewSimpleTransaction(from, to string, amount int64) *Transaction {
	//1.定义Input和Output的数组
	var txInputs []*TxInput
	var txOutputs []*TxOutput

	//2.创建Input,暂时写固定数值
	idBytes, _ := hex.DecodeString("9d55992f32659a797871a0536961ff84c406b6fb69f2d79947a377d3542582ee")

	txInput := &TxInput{TxID: idBytes, Vout: 1, ScriptSiq: from}
	txInputs = append(txInputs, txInput)

	//3.创建OutPut
	//转账
	txOutput := &TxOutput{Value: amount, ScriptPubKey: to}
	txOutputs = append(txOutputs, txOutput)
	//找零
	txOutput2 := &TxOutput{Value: 7 - amount, ScriptPubKey: from}
	txOutputs = append(txOutputs, txOutput2)

	//4.创建交易
	tx := &Transaction{TxID: []byte{}, Vins: txInputs, Vouts: txOutputs}

	//5.设置交易ID
	tx.SetID()
	return tx

}

//判断tx是否为CoinBase交易
func (tx *Transaction) IsCoinBaseTransaction() bool {
	return len(tx.Vins[0].TxID) == 0 && tx.Vins[0].Vout == -1

}
