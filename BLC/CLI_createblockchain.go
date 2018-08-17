package BLC

import (
	"fmt"
	"os"
)

func (cli *CLI) CreateBlockChain(address string) {
	//fmt.Println("创世区块。。。")
	CreateBlockChainWithGenesisBlock(address)

	//重置：
	bc :=GetBlockChainObject()
	if bc == nil{
		fmt.Println("没有数据库。。")
		os.Exit(1)
	}
	defer bc.DB.Close()
	utxoSet:=&UTXOSet{bc}
	utxoSet.ResetUTXOSet()
}
