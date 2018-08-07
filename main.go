package main

import (
	"MyPublicChain/BLC"
	"fmt"
)

func main() {
	//1.创建一个带有创世区块的区块链
	bc := BLC.CreateBlockChainWithGenesisBlock("Create Block Chain With GenesisBlock")
	//2.通过POW挖出第一个区块
	bc.AddBlockToBlockChain("Send $B100 To brucefeng", bc.Blocks[len(bc.Blocks)-1].Hash, bc.Blocks[len(bc.Blocks)-1].Height+1)
	//3.验证区块是否合法
	pow := BLC.NewProofOfWork(bc.Blocks[1])
	isValid := pow.IsValid()
	if isValid {
		fmt.Println("区块合法")
		return
	}
	fmt.Println("区块不合法")

}
