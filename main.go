package main

import (
	"MyPublicChain/BLC"
)

func main() {

	//添加新的区块至DB中

	bc := BLC.CreateBlockChainWithGenesisBlock("brucefeng GenesisBloc")


	bc.AddBlockToBlockChain("New")
	bc.PrintChains()
	defer bc.DB.Close()



}
