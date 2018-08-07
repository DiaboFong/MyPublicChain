package main

import (
	"MyPublicChain/BLC"
)

func main() {
	//1.创建一个带有创世区块的区块链
	bc := BLC.CreateBlockChainWithGenesisBlock("Create Block Chain With GenesisBlock")
	//2.通过POW挖出第一个区块
	bc.AddBlockToBlockChain("Send $B100 To brucefeng", bc.Blocks[len(bc.Blocks)-1].Hash, bc.Blocks[len(bc.Blocks)-1].Height+1)

}
