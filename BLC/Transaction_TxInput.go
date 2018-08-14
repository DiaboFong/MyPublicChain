package BLC

type TxInput struct {
	//1.交易ID：引用的TxOutput所在的交易ID
	TxID []byte

	//2.引用的交易中的哪个txoutput,其实就是下标
	Vout int

	//3.输入脚本，也就是解锁脚本。暂时理解为用户名
	ScriptSiq string
}

//判断TxInput是否时指定的用户消费
func (txInput *TxInput) UnlockWithAddress(address string) bool{
	return txInput.ScriptSiq == address
}