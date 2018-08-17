package BLC

import (
	"fmt"
	"os"
)

func (cli *CLI) GetBalance(address string) {
	bc := GetBlockChainObject()
	if bc == nil {
		fmt.Println("没有BlockChain，无法查询。。")
		os.Exit(1)
	}
	defer bc.DB.Close()
	//total := bc.GetBalance(address,[]*Transaction{})
	utxoSet := &UTXOSet{bc}
	total := utxoSet.GetBalance(address)

	fmt.Printf("%s,余额是：%d\n", address, total)
}
