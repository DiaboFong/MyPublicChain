package BLC

//UTXO:Unspent Transaction output
type UTXO struct {
	//1.该output所在的交易id
	TxID []byte
	//2.该output 的下标
	Index int
	//3.output,未花费的
	Output *TxOutput
}
