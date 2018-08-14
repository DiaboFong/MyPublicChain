package BLC

import (
	"fmt"
	"os"
)

func (cli *CLI) CreateBlockChain(address string) {
	//fmt.Println("创世区块。。。")
	CreateBlockChainWithGenesisBlock(address)

}
