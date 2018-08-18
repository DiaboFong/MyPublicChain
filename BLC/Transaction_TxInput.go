package BLC

import "bytes"

type TxInput struct {
	//1.交易ID：引用的TxOutput所在的交易ID
	TxID []byte

	//2.引用的交易中的哪个txoutput,其实就是下标
	Vout int

	//3.输入脚本，也就是解锁脚本。暂时理解为用户名
	//ScriptSiq string

	Signature []byte //数字签名
	PublicKey[]byte //原始公钥，钱包里的公钥

}

//判断TxInput是否时指定的用户消费
func (txInput *TxInput) UnlockWithAddress(pubKeyHash []byte) bool{

	pubKeyHash2:=PubKeyHash(txInput.PublicKey)
	return bytes.Compare(pubKeyHash,pubKeyHash2) == 0
}