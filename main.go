package main

import (
	"fmt"
	"MyPublicChain/BLC"
)

func main() {
	//创建一个带有创世区块的区块链
	bc := BLC.CreateBlockChainWithGenesisBlock("Create Block Chain With GenesisBlock")
	//打印该区块链，&{[0xc420068060]}
	fmt.Println(bc)
	//打印该区块链的Blocks
	fmt.Println(bc.Blocks)
	//打印该区块链Blocks数组中存储的第一个区块
	fmt.Println(bc.Blocks[0])

}
