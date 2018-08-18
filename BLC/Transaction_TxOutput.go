package BLC

import "bytes"

//定义TxOutput的结构体
type TxOutput struct {
	//金额
	Value int64  //金额
	//锁定脚本，也叫输出脚本，公钥，目前先理解为用户名，钥花费这笔前，必须钥先解锁脚本
	//ScriptPubKey string
	PubKeyHash [] byte//公钥哈希
}

//判断TxOutput是否时指定的用户解锁
func (txOutput *TxOutput) UnlockWithAddress(address string) bool{
	full_payload:=Base58Decode([]byte(address))

	pubKeyHash:=full_payload[1:len(full_payload)-addressCheckSumLen]

	return bytes.Compare(pubKeyHash,txOutput.PubKeyHash) == 0
}

//根据地址创建一个output对象
func NewTxOutput(value int64,address string) *TxOutput{
	txOutput:=&TxOutput{value,nil}
	txOutput.Lock(address)
	return txOutput
}

//锁定
func (tx *TxOutput) Lock(address string){
	full_payload := Base58Decode([]byte(address))
	tx.PubKeyHash = full_payload[1:len(full_payload)-addressCheckSumLen]
}

